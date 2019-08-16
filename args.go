package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var verbose = false

// Args struct
type Args struct {
	Dir  string
	In   string
	Out  string
	Save string
}

// ParseFlags ...
func ParseFlags() *Args {
	dir := flag.String("dir", ".", "The directory location of your files")
	in := flag.String("in", ".png", "The filetype to start with")
	out := flag.String("out", ".jpg", "The filetype to convert to")
	saveLoc := flag.String("save", ".", "The save location for files")
	help := flag.Bool("help", false, "Print flags")
	verboseLogs := flag.Bool("v", false, "This option turns on verbose logging")

	flag.Parse()

	if *help {
		logger.I("USAGE")
		flag.PrintDefaults()
		os.Exit(0)
	}

	verbose = *verboseLogs
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
		fmt.Printf("The following error occured parsing flags: %v", err.Error())
		os.Exit(1)
	}

	return path
}
