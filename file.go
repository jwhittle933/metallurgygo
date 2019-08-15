package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// File ...
type File struct {
	Path    string        // Absolute path location to file
	FromExt string        // Ext for checking filetype and skipping if not --in type
	Name    string        // File name
	ToExt   string        // File type to convert to
	OutPath string        // OutPath absolute path for converted image
	Data    image.Image   // Data for holding decoded image
	Buffer  *bytes.Buffer // Buffer for encoded data
}

// Decode ...
func (inf *File) Decode() {
	f, err := os.Open(inf.Path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var img image.Image
	switch filepath.Ext(inf.Path) {
	case ".jpg", ".jpeg", ".JPEG":
		img, err = jpeg.Decode(f)
	case ".png":
		img, err = png.Decode(f)
	}

	if err != nil {
		panic(err)
	}

	inf.Data = img
}

// Encode ...
func (inf *File) Encode() {
	var b []byte
	buf := bytes.NewBuffer(b)

	var err error
	switch inf.ToExt {
	case ".jpg", ".jpeg", ".JPEG":
		err = jpeg.Encode(buf, inf.Data, nil)
	case ".png":
		err = png.Encode(buf, inf.Data)
	}

	if err != nil {
		fmt.Errorf("Encode error %s", err.Error())
	}

	inf.Buffer = buf
}
