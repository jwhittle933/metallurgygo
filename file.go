package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"fmt"
	"os"
	"path/filepath"
)

// File ...
type File struct {
	// Absolute path location to file
	Path string
	// Ext for checking filetype and skipping if not --in type
	FromExt string
	// File name
	Name string
	// File type to convert to
	ToExt string
	// OutPath absolute path for converted image
	OutPath string
	// Data for holding decoded image
	Data image.Image
	// Buffer for encoded data
	Buffer *bytes.Buffer
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
