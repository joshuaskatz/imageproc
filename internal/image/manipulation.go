package internal_image

import "image"

func (file *File) WhiteBalance() {
	file.Image = image.Black.Bounds()
}
