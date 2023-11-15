package internal_image

import (
	"image"
	"log"
	"os"

	"golang.org/x/image/tiff"
)

type File struct {
	Image         image.Image
	Path          string
	InvertedImage image.NRGBA
}

func NewFile(filepath string) *File {
	return &File{Path: filepath}
}

func (file *File) LoadImage() {
	f, err := os.Open(file.Path)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()

	i, err := tiff.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}

	file.Image = i
}

func (file *File) SaveImage() {
	f, err := os.Create("/Users/joshkatz/Desktop/example.tif")

	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()

	err = tiff.Encode(f, &file.InvertedImage, &tiff.Options{})
	if err != nil {
		log.Fatalf("tiff.Encode failed: %v", err)
	}
}
