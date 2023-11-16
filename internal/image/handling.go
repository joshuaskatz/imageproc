package internal_image

import (
	"image"
	"log"
	"os"

	"github.com/chai2010/tiff"
)

type File struct {
	InitImage     image.Image
	Path          string
	InvertedImage image.Image
	IsColor       bool
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

	file.InitImage = i
}

func (file *File) SaveImage() {
	f, err := os.Create("/Users/joshkatz/Desktop/example.tif")

	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()

	err = tiff.Encode(f, file.InvertedImage, &tiff.Options{})
	if err != nil {
		log.Fatalf("tiff.Encode failed: %v", err)
	}
}
