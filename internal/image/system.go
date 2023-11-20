package image

import (
	"log"
	"os"

	"github.com/chai2010/tiff"
)

func (i *ImageFile) SaveImage() {
	f, err := os.Create("/Users/joshkatz/Desktop/example.tif")

	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()

	err = tiff.Encode(f, i.ConvertedImage, &tiff.Options{})
	if err != nil {
		log.Fatalf("tiff.Encode failed: %v", err)
	}
}
