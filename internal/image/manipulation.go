package internal_image

import "image"

func (file *File) WhiteBalance() {
	file.InitImage = image.Black.Bounds()
}
