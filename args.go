package main

import (
	"flag"
)

// Args struct
type Args struct {
	Dir string
	In string
	Out string
}

// ParseFlags ...
func ParseFlags() *Args {
	dir := flag.String("dir", ".", "The directory location of your files")
	in := flag.String("in", ".png", "The filetype to start with")
	out := flag.String("out", ".jpg", "The filetype to convert to")

	flag.Parse()

	return &Args {
		Dir: *dir,
		In: *in,
		Out: *out,
	}
}
