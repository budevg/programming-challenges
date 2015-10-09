package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
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

	for remainingDecodeSize > 0 {
		count, err := decodeNextTrack(f, &p)
		if err != nil {
			return nil, err
		}
		remainingDecodeSize -= count
	}

	return &p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	version [32]byte
	tempo   float32
	tracks  []Track
}

type Track struct {
	id    uint32
	name  string
	steps [16]bool
}

func (p *Pattern) String() string {
	res := []string{}
	res = append(res, fmt.Sprintf("Saved with HW Version: %s", strings.TrimRight(string(p.version[:]), "\x00")))
	res = append(res, fmt.Sprintf("Tempo: %s", toPrettyFloat(p.tempo)))
	for _, track := range p.tracks {
		steps := bytes.Buffer{}
		for i := 0; i < 16; i++ {
			if i%4 == 0 {
				steps.WriteString("|")
			}
			if track.steps[i] {
				steps.WriteString("x")
			} else {
				steps.WriteString("-")
			}
		}
		steps.WriteString("|")

		res = append(res, fmt.Sprintf("(%d) %s\t%s", track.id, track.name, steps.String()))
	}
	res = append(res, "")
	return strings.Join(res, "\n")
}

func toPrettyFloat(f float32) string {
	if float32(int(f)) == f {
		return fmt.Sprintf("%d", int(f))
	} else {
		return fmt.Sprintf("%.1f", f)
	}
}

func decodeHeader(f *os.File, p *Pattern) (uint64, error) {
	magic := make([]byte, 6)
	err := binary.Read(f, binary.LittleEndian, &magic)
	if err != nil {
		return 0, err
	}

	if string(magic) != "SPLICE" {
		return 0, fmt.Errorf("invalid magic [%s]", string(magic))
	}

	var encodingSize uint64
	err = binary.Read(f, binary.BigEndian, &encodingSize)
	if err != nil {
		return 0, err
	}

	err = binary.Read(f, binary.BigEndian, &p.version)
	if err != nil {
		return 0, err
	}
	encodingSize -= 32

	err = binary.Read(f, binary.LittleEndian, &p.tempo)
	if err != nil {
		return 0, err
	}
	encodingSize -= 4

	return encodingSize, nil
}

func decodeNextTrack(f *os.File, p *Pattern) (uint64, error) {
	t := Track{}
	count := 0

	err := binary.Read(f, binary.LittleEndian, &t.id)
	if err != nil {
		return 0, err
	}
	count += 4

	var nameSize byte
	err = binary.Read(f, binary.LittleEndian, &nameSize)
	if err != nil {
		return 0, err
	}
	count += 1

	name := make([]byte, nameSize)
	err = binary.Read(f, binary.LittleEndian, &name)
	if err != nil {
		return 0, err
	}
	count += int(nameSize)

	t.name = string(name)
	for i := 0; i < 16; i++ {
		var b byte
		err = binary.Read(f, binary.LittleEndian, &b)
		if err != nil {
			return 0, err
		}
		t.steps[i] = (b == 1)

	}
	count += 16

	p.tracks = append(p.tracks, t)
	return uint64(count), nil

}
