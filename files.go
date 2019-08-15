package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
		file, err := os.Create(filepath.Join(f.OutPath, f.Name)+f.ToExt)
		if err != nil {
			fmt.Errorf("Write error: %s", err.Error())
		}
		defer file.Close()

		_, err = file.Write(f.Buffer.Bytes())

	}
}

func must(files []os.FileInfo, err error) []os.FileInfo {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return files
}

func pathAndExt(dir string, file os.FileInfo) (string, string) {
	path := dir+"/"+file.Name()
	ext := filepath.Ext(path)
	return path, ext
}
