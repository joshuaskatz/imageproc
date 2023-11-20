package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func CanvasInitImage(w fyne.Window, i *image.ImageFile) *canvas.Image {
	initImage := canvas.NewImageFromImage(i.InitImage)

	initImage.FillMode = canvas.ImageFillContain
	return initImage
}

func CanvasConvertedImage(w fyne.Window, i *image.ImageFile) *canvas.Image {
	convertedImage := canvas.NewImageFromImage(i.ConvertedImage)

	convertedImage.FillMode = canvas.ImageFillContain
	return convertedImage
}
