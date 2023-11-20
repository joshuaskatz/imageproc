package image

import (
	"fmt"
	"image"

	"github.com/davidbyttow/govips/v2/vips"
)

type ImageFile struct {
	InitImage      image.Image
	Path           string
	ConvertedImage image.Image
	IsColor        bool
	vipsRef        *vips.ImageRef
}

func NewFile(filePath string) *ImageFile {
	i, err := vips.NewImageFromFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	exportParams := vips.NewDefaultExportParams()

	initImage, err := i.ToImage(exportParams)

	return &ImageFile{vipsRef: i, InitImage: initImage}
}

func (file *ImageFile) SetIsColor(isColor bool) {
	file.IsColor = isColor
}

func (file *ImageFile) SetFilePath(filePath string) {
	file.Path = filePath
}
