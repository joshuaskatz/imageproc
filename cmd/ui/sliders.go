package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/joshuaskatz/imageproc/internal/image"
)

func Sliders(w fyne.Window, i *image.ImageFile, canvasImage *canvas.Image) (*widget.Slider, *widget.Slider, *widget.Slider) {
	rBinding := binding.NewFloat()
	gBinding := binding.NewFloat()
	bBinding := binding.NewFloat()

	rSlider := widget.NewSliderWithData(-100, 100, rBinding)
	gSlider := widget.NewSliderWithData(-100, 100, gBinding)
	bSlider := widget.NewSliderWithData(-100, 100, bBinding)

	// i.RGBChannels(rBinding, gBinding, bBinding, func() { w.Canvas().Refresh(canvasImage) })

	return rSlider, gSlider, bSlider
}
