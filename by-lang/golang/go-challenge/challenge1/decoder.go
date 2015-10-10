package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"io"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Pattern{}

	remainingDecodeSize, err := decodeHeader(f, &p)
	if err != nil {
		return nil, err
	}

	r := io.LimitReader(f, int64(remainingDecodeSize))
	for {
		err = decodeNextTrack(r, &p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
}

func (p *Pattern) String() string {
	var res []string
	res = append(res, fmt.Sprintf("Saved with HW Version: %s", strings.TrimRight(string(p.version[:]), "\x00")))
	res = append(res, fmt.Sprintf("Tempo: %s", toPrettyFloat(p.tempo)))
	for _, t := range p.tracks {
		steps := bytes.Buffer{}
		for i := 0; i < 16; i++ {
			if i%4 == 0 {
				steps.WriteString("|")
			}
			if t.steps[i] {
				steps.WriteString("x")
			} else {
				steps.WriteString("-")
			}
		}
		steps.WriteString("|")

		res = append(res, fmt.Sprintf("(%d) %s\t%s", t.id, t.name, steps.String()))
	}
	res = append(res, "")
	return strings.Join(res, "\n")
}

func toPrettyFloat(f float32) string {
	if float32(int(f)) == f {
		return fmt.Sprintf("%d", int(f))
	}

	return fmt.Sprintf("%.1f", f)
}

type errReader struct {
	r io.Reader
	err error
}

func (er *errReader) Read(order binary.ByteOrder, data interface{}) {
	if er.err != nil {
		return
	}

	er.err = binary.Read(er.r, order, data)
}

func decodeHeader(r io.Reader, p *Pattern) (uint64, error) {
	er := &errReader{r : r}

	magic := make([]byte, 6)
	er.Read(binary.LittleEndian, &magic)
	if er.err == nil && string(magic) != "SPLICE" {
		return 0, fmt.Errorf("invalid magic [%s]", string(magic))
	}

	var encodingSize uint64
	er.Read(binary.BigEndian, &encodingSize)

	er.Read(binary.BigEndian, &p.version)
	encodingSize -= 32

	er.Read(binary.LittleEndian, &p.tempo)
	encodingSize -= 4

	if er.err != nil {
		return  0, er.err
	}
	return encodingSize, nil
}

func decodeNextTrack(r io.Reader, p *Pattern) error {
	er := &errReader{r : r}
	t := Track{}

	er.Read(binary.LittleEndian, &t.id)

	var nameSize byte
	er.Read(binary.LittleEndian, &nameSize)

	name := make([]byte, nameSize)
	er.Read(binary.LittleEndian, &name)

	t.name = string(name)
	for i := 0; i < 16; i++ {
		var b byte
		er.Read(binary.LittleEndian, &b)
		t.steps[i] = (b == 1)

	}

	if er.err != nil {
		return er.err
	}

	p.tracks = append(p.tracks, t)
	return nil
}
