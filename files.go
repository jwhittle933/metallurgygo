package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// InFile ...
type InFile struct {
	// Path for passing to jpg.Decode
	Path string
	// Ext for checking filetype and skipping
	// if not --in type
	Ext  string
}

// InFiles ...
type InFiles []InFile


// GetFiles ...
func GetFiles(dir string) InFiles {
	files := Must(ioutil.ReadDir(dir))

	inFiles := InFiles{}
	for _, f := range files {
		path, ext := pathAndExt(dir, f)

		i := InFile{
			Path: path,
			Ext:  ext,
		}

		inFiles = append(inFiles, i)
	}

	return inFiles
}

// Filter ...
func (i InFiles) Filter(ext string) InFiles {
	inFiles := InFiles{}
	for _, f := range i {
		if f.Ext == ext {
			inFiles = append(inFiles, f)
		}
	}

	return inFiles
}

// Must ...
func Must(files []os.FileInfo, err error) []os.FileInfo {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return files
}


func pathAndExt(dir string, file os.FileInfo) (string, string) {
	path := filepath.Join(dir, file.Name())
	ext := filepath.Ext(path)
	return path, ext
}
