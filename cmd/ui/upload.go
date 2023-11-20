package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func ShowUploadButton(w fyne.Window, i *image.ImageFile) {
	uploadButton := UploadButton(w, i)

	uploadButtonContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), uploadButton, layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), uploadButtonContainer))

}
