package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Files ...
type Files []*File

// NewFiles ...
func NewFiles(args *Args) Files {
	files := must(ioutil.ReadDir(args.Dir))

	var fs []*File
	for _, f := range files {
		path, ext := pathAndExt(args.Dir, f)
		// if file not with --in ext, skip
		if ext != args.In {
			continue
		}

		if verbose {
			logger.I("Found %s", f.Name())
		}

		fs = append(fs, &File{
			Path:    path,
			FromExt: ext,
			Name:    f.Name()[:len(f.Name())-len(ext)],
			ToExt:   args.Out,
			OutPath: args.Save,
		})
	}

	// Check for zero results
	if len(fs) < 1 {
		logger.E("\nSearch returned no results matching extension \"%s\" in %s\n\n", args.In, args.Dir)
		os.Exit(0)
	}

	if !verbose {
		logger.I("Found %v files\n\n", len(fs))
	}

	return Files(fs)
}

// Decode ...
func (infs Files) Decode() Files {
	for _, file := range infs {
		file.Decode()
	}

	return infs
}

// Encode ...
func (infs Files) Encode() Files {
	for _, file := range infs {
		file.Encode()
	}

	return infs
}

// Write ...
func (infs Files) Write() {
	for _, f := range infs {
		f.Write()
	}
}

func must(files []os.FileInfo, err error) []os.FileInfo {
	if err != nil {
		fmt.Printf("The following error occured: %v", err.Error())
		os.Exit(1)
	}

	return files
}

func pathAndExt(dir string, file os.FileInfo) (string, string) {
	path := dir+"/"+file.Name()
	ext := filepath.Ext(path)
	return path, ext
}
