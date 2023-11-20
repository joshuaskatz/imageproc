package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func ShowConvertView(w fyne.Window, i *image.ImageFile) {
	initImage := CanvasInitImage(w, i)

	check := widget.NewCheck("Is Color", func(value bool) {
		i.SetIsColor(value)
	})

	convertButton := ConvertButton(w, i)

	mainViewContainer := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(), initImage, layout.NewSpacer(), check, layout.NewSpacer(), convertButton)
	w.SetContent(container.New(layout.NewVBoxLayout(), mainViewContainer))
}
