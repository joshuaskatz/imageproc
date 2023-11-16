package image

import (
	"log"
	"os"

	"github.com/chai2010/tiff"
)

func (file *ImageFile) LoadImage() {
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

func (file *ImageFile) SaveImage() {
	f, err := os.Create("/Users/joshkatz/Desktop/example.tif")

	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()

	err = tiff.Encode(f, file.ConvertedImage, &tiff.Options{})
	if err != nil {
		log.Fatalf("tiff.Encode failed: %v", err)
	}
}
