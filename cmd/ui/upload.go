package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ShowUploadButton(w fyne.Window) {
	uploadButton := UploadButton(w)

	uploadButtonContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), uploadButton, layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), uploadButtonContainer))

}
