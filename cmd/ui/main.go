package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()
	myApp := app.New()
	w := myApp.NewWindow("Button Widget")

	ShowUploadButton(w)

	w.ShowAndRun()
}
