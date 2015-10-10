// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	version [32]byte
	tempo   float32
	tracks  []Track
}

// Track represents single track details
type Track struct {
	id    uint32
	name  string
	steps [16]bool
}
