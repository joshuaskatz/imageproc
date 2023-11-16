package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/joshuaskatz/imageproc/internal/image"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")
	var uri string
	content := widget.NewButton("click me", func() {
		d := dialog.NewFileOpen(func(f fyne.URIReadCloser, e error) {
			uri = f.URI().Path()

			file := image.NewFile()

			file.SetFilePath(uri)

			file.LoadImage()

			file.SetIsColor(true)

			file.Conversion()

			file.SaveImage()

		}, myWindow)

		d.Show()
	})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
