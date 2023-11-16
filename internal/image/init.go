package image

import "image"

type ImageFile struct {
	InitImage      image.Image
	Path           string
	ConvertedImage image.Image
	IsColor        bool
}

func NewFile() *ImageFile {
	return &ImageFile{}
}

func (file *ImageFile) SetIsColor(isColor bool) {
	file.IsColor = isColor
}

func (file *ImageFile) SetFilePath(filePath string) {
	file.Path = filePath
}
