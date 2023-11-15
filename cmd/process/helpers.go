package main

import (
	"image"

	internal_image "github.com/joshuaskatz/imageproc/internal/image"

	"github.com/disintegration/gift"
)

func Process(file *internal_image.File) *image.NRGBA {
	invert := gift.Invert()
	gi := gift.New(invert)
	invertDst := image.NewNRGBA(gi.Bounds(file.Image.Bounds()))
	gi.Draw(invertDst, file.Image)

	colorBalance := gift.ColorFunc(
		func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
			r = r0 + 0.15
			g = g0
			b = b0 - 0.05
			a = a0
			return r, g, b, a
		},
	)
	gcb := gift.New(colorBalance)
	colorBalanceDst := image.NewNRGBA(gcb.Bounds(invertDst.Rect))
	gcb.Draw(colorBalanceDst, invertDst)

	contrast := gift.Contrast(10)
	gc := gift.New(contrast)
	contrastDst := image.NewNRGBA(gc.Bounds(colorBalanceDst.Rect))
	gc.Draw(contrastDst, colorBalanceDst)

	brightness := gift.Gamma(0.15)
	gb := gift.New(brightness)
	brightnessDst := image.NewNRGBA(gb.Bounds(contrastDst.Rect))
	gb.Draw(brightnessDst, contrastDst)

	// break out later
	return brightnessDst
}
