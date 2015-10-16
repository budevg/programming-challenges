package main

import (
        "flag"
        "fmt"
        "io"
        "log"
        "net"
        "os"
	"golang.org/x/crypto/nacl/box"
	"encoding/binary"
	"crypto/rand"
)

type SecureReader struct {
	r io.Reader
	sharedKey *[32]byte
}

func (sr SecureReader) Read(p []byte) (int, error) {
	var nonce [24]byte

	_, err := io.ReadFull(sr.r, nonce[:])
	if err != nil {
		return 0, err
	}

	var msgSize uint32
	err = binary.Read(sr.r, binary.LittleEndian, &msgSize)
	if err != nil {
		return 0, err
	}

	// msg := box.SealAfterPrecomputation(nil, p, nonce, sw.sharedKey)
	msg := make([]byte, msgSize)
	_, err = io.ReadFull(sr.r, msg)
	if err != nil {
		return 0, err
	}

	decodedMsg, ok := box.OpenAfterPrecomputation(nil, msg, &nonce, sr.sharedKey)
	if !ok {
		return 0, fmt.Errorf("Failed to decode")
	}
	n := copy(p, decodedMsg)
	return n, nil

}
// NewSecureReader instantiates a new SecureReader
func NewSecureReader(r io.Reader, priv, pub *[32]byte) io.Reader {
	sr := SecureReader{
		r : r,
		sharedKey : &[32]byte{},
	}
	box.Precompute(sr.sharedKey, pub, priv)
        return sr
}

type SecureWriter struct {
	w io.Writer
	sharedKey *[32]byte
}

func generateNonce() (*[24]byte, error) {
	nonce := [24]byte{}
	_, err := rand.Read(nonce[:])
	if err != nil {
		return nil, err
	}

	return &nonce, nil
}

func (sw SecureWriter) Write(p []byte) (int, error) {
	nonce, err := generateNonce()
	if err != nil {
		return 0, err
	}

	_, err = sw.w.Write(nonce[:])
	if err != nil {
		return 0, err
	}

	msgSize := uint32(len(p) + box.Overhead)
	err = binary.Write(sw.w, binary.LittleEndian, &msgSize)
	if err != nil {
		return 0, err
	}

	msg := box.SealAfterPrecomputation(nil, p, nonce, sw.sharedKey)
	n, err := sw.w.Write(msg)
	if err != nil {
		return 0, err
	}

	n -= box.Overhead
	return n, nil
}

// NewSecureWriter instantiates a new SecureWriter
func NewSecureWriter(w io.Writer, priv, pub *[32]byte) io.Writer {
	sw := SecureWriter{
		w : w,
		sharedKey : &[32]byte{},
	}
	box.Precompute(sw.sharedKey, pub, priv)
        return sw
}

type SecureReadWriteCloser struct {
	io.Reader
	io.Writer
	io.Closer
}
// Dial generates a private/public key pair,
// connects to the server, perform the handshake
// and return a reader/writer.
func Dial(addr string) (io.ReadWriteCloser, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(pub[:])
	if err != nil {
		return nil, err
	}

	var peerPub [32]byte
	_, err = io.ReadFull(conn, peerPub[:])
	if err != nil {
		return nil, err
	}

        return SecureReadWriteCloser{
		NewSecureReader(conn, priv, &peerPub),
		NewSecureWriter(conn, priv, &peerPub),
		conn,
	}, nil
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return
	}

	_, err = conn.Write(pub[:])
	if err != nil {
		return
	}

	var peerPub [32]byte
	_, err = io.ReadFull(conn, peerPub[:])
	if err != nil {
		return
	}

	sr := NewSecureReader(conn, priv, &peerPub)
	sw := NewSecureWriter(conn, priv, &peerPub)
	for {
		var buffer [128]byte
		n, err := sr.Read(buffer[:])
		if err != nil {
			return
		}
		sw.Write(buffer[:n])
	}

}

// Serve starts a secure echo server on the given listener.
func Serve(l net.Listener) error {
        for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go HandleConn(conn)
	}
}

func main() {
        port := flag.Int("l", 0, "Listen mode. Specify port")
        flag.Parse()

        // Server mode
        if *port != 0 {
                l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
                if err != nil {
                        log.Fatal(err)
                }
                defer l.Close()
                log.Fatal(Serve(l))
        }

        // Client mode
        if len(os.Args) != 3 {
                log.Fatalf("Usage: %s <port> <message>", os.Args[0])
        }
        conn, err := Dial("localhost:" + os.Args[1])
        if err != nil {
                log.Fatal(err)
        }
        if _, err := conn.Write([]byte(os.Args[2])); err != nil {
                log.Fatal(err)
        }
        buf := make([]byte, len(os.Args[2]))
        n, err := conn.Read(buf)
        if err != nil && err != io.EOF {
                log.Fatal(err)
        }
        fmt.Printf("%s\n", buf[:n])
}
