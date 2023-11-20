package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Button Widget")
	file := image.NewFile()

	ShowUploadButton(w, file)

	w.ShowAndRun()
}
