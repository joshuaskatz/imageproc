package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func ShowEditView(w fyne.Window, i *image.ImageFile) {
	convertedImage := CanvasConvertedImage(w, i)

	// rSlider, gSlider, bSlider := Sliders(w, i, convertedImage)

	mainViewContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), convertedImage, layout.NewSpacer(), layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), mainViewContainer))

}
