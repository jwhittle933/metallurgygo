package main

import (
	"flag"
	"path/filepath"
)

// Args struct
type Args struct {
	Dir  string
	In   string
	Out  string
	Save string
}

// ParseFlags ...
func ParseFlags() *Args {
	dir := flag.String("dir", ".", "The directory location of your files. DEFAULT: ./")
	in := flag.String("in", ".png", "The filetype to start with. DEFAULT: .png")
	out := flag.String("out", ".jpg", "The filetype to convert to. DEFAULT: .jpg")
	saveLoc := flag.String("save", ".", "The save location for files. DEFAULT: ./")

	flag.Parse()

	return &Args{
		Dir:  Absolute(dir),
		In:   *in,
		Out:  *out,
		Save: Absolute(saveLoc),
	}
}

// Absolute ...
func Absolute(dir *string) string {
	path, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}

	return path
}
