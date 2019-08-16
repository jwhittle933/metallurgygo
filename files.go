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

	var paths []*File
	for _, f := range files {
		path, ext := pathAndExt(args.Dir, f)
		// if file not with --in ext, skip
		if ext != args.In {
			continue
		}

		paths = append(paths, &File{
			Path:    path,
			FromExt: ext,
			Name:    f.Name()[:len(f.Name())-len(ext)],
			ToExt:   args.Out,
			OutPath: args.Save,
		})
	}

	// Check for zero results
	if len(paths) < 1 {
		logger.E("\nSearch returned no results matching extension \"%s\" in %s\n\n", args.In, args.Dir)
		os.Exit(0)
	}

	return Files(paths)
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
	for _, f := range infs {
		f.Encode()
	}

	return infs
}

// Write ...
func (infs Files) Write() {
	for _, f := range infs {
		if len(f.Buffer.Bytes()) < 1 {
			continue
		}
		file, err := os.Create(filepath.Join(f.OutPath, f.Name)+f.ToExt)
		if err != nil {
			logger.E("Write error for %s: %s", f.Name, err.Error())
			file.Close()
			continue
		}

		_, err = file.Write(f.Buffer.Bytes())
		if err != nil {
			logger.E("Write error for %s: %s", f.Name, err.Error())
			file.Close()
			continue
		}
		file.Close()
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
