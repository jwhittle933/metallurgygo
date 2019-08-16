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
		defer Close(f)()
		return
	}

	img, err := inf.decoder(f)

	if err != nil {
		logger.E("\nError decoding %s: %s", inf.Path, err.Error())
		defer Close(f)()
		return
	}

	if verbose {
		logger.I("Decoding %s", inf.Name+inf.FromExt)
	}

	inf.Data = img
	defer Close(f)()
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

// Write ...
func (inf *File) Write() {
	// In the case of encoding error, Buffer will be empty
	// Pass to next image if buffer is empty
	if len(inf.Buffer.Bytes()) < 1 {
		logger.E("Skipping %s due to encoding error\n", inf.Name+inf.FromExt)
		return
	}
	file, err := os.Create(filepath.Join(inf.OutPath, inf.Name)+inf.ToExt)
	if err != nil {
		logger.E("Write error for %s: %s", inf.Name, err.Error())
		defer Close(file)()
		return
	}

	_, err = file.Write(inf.Buffer.Bytes())
	if err != nil {
		logger.E("Write error for %s: %s", inf.Name, err.Error())
		defer Close(file)()
		return
	}

	if verbose {
		logger.I("Wrote %s", inf.Name+inf.ToExt)
	}

	defer Close(file)()
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

// Close ...
func Close(f *os.File) func() {
	return func() {
		err := f.Close()
		if err != nil {
			logger.E("Error closing file: %s", err.Error())
		}
	}
}
