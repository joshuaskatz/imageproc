package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func UploadButton(w fyne.Window, i *image.ImageFile) *widget.Button {
	var uri string

	button := widget.NewButton("Upload Image", func() {
		d := dialog.NewFileOpen(func(f fyne.URIReadCloser, e error) {
			uri = f.URI().Path()

			i.SetFilePath(uri)
			i.LoadImage()

			ShowConvertView(w, i)
		}, w)

		d.Show()
	})

	return button
}

func ConvertButton(w fyne.Window, i *image.ImageFile) *widget.Button {

	button := widget.NewButton("Convert Image", func() {
		i.Conversion()

		ShowEditView(w, i)
	})

	return button
}
