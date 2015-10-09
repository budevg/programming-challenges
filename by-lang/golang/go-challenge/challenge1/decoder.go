package drum

import (
	"fmt"
	"strings"
	"bytes"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	p := &Pattern{
		version : "0.808-alpha",
		tempo : 120,
		tracks : []struct {
			id int
			name string
			steps [16]bool
		}{
			{
				id : 0,
				name : "kick",
				steps : [16]bool{
					true,
					false,
					false,
					false,
					true,
					false,
					false,
					false,
					true,
					false,
					false,
					false,
					true,
					false,
					false,
					false,
				},
			},
		},
	}
	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo float32
	tracks []struct {
		id int
		name string
		steps [16]bool
	}
}

func (p *Pattern) String()  string {
	res := []string{}
	res = append(res, fmt.Sprintf("Saved with HW Version: %s", p.version))
	res = append(res, fmt.Sprintf("Tempo: %s", toPrettyFloat(p.tempo)))
	for _, track := range p.tracks {
		steps := bytes.Buffer{}
		for i := 0; i < 16; i++ {
			if i % 4 == 0 {
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
	return strings.Join(res, "\n")
}


func toPrettyFloat(f float32) string {
	if float32(int(f)) == f {
		return fmt.Sprintf("%d", int(f))
	} else {
		return fmt.Sprintf("%.2f", f)
	}
}
