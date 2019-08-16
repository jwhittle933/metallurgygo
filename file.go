package main

import (
	"bytes"
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
		logger.E("\nError opening %s: %s", inf.Path, err.Error())
		f.Close()
		return
	}

	img, err := inf.decoder(f)

	if err != nil {
		logger.E("\nError decoding %s: %s", inf.Path, err.Error())
		f.Close()
		return
	}

	if verbose {
		logger.I("Decoding %s", inf.Name+inf.FromExt)
	}

	inf.Data = img
	f.Close()
}

// Encode ...
func (inf *File) Encode() {
	var b []byte
	buf := bytes.NewBuffer(b)
	err := inf.encoder(buf)

	if err != nil {
		logger.E("\nError encoding %s: %s", inf.Path, err.Error())
		return
	}

	if verbose {
		logger.I("Encoding %s", inf.Name+inf.ToExt)
	}

	inf.Buffer = buf
}

func (inf *File) decoder(f *os.File) (image.Image, error) {
	var img image.Image
	var err error
	switch filepath.Ext(inf.Path) {
	case ".jpg", ".jpeg", ".JPEG":
		img, err = jpeg.Decode(f)
	case ".png":
		img, err = png.Decode(f)
	}

	return img, err
}

func (inf *File) encoder(buf *bytes.Buffer) error {
	var err error
	switch inf.ToExt {
	case ".jpg", ".jpeg", ".JPEG":
		// TODO: extend cli to include image quality
		err = jpeg.Encode(buf, inf.Data, &jpeg.Options{
			Quality: 50,
		})
	case ".png":
		err = png.Encode(buf, inf.Data)
	}

	return err
}
