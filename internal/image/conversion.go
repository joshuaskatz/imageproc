package image

import (
	"image"

	"github.com/disintegration/gift"
)

// func (file *ImageFile) Conversion() {
// 	invertedImage := effect.Invert(file.InitImage)
// 	colorCorrectedImage := invertedImage
// 	if file.IsColor {
// 		colorCorrectedImage = adjust.Apply(invertedImage, func(c color.RGBA) color.RGBA {
// 			nrgba := color.NRGBA{c.R + 40, c.G, c.B - 15, c.A}

// 			r, g, b, a := nrgba.RGBA()

// 			return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
// 		})
// 	}
// 	gammaCorrectedImage := adjust.Gamma(colorCorrectedImage, 0.25)
// 	contrastCorrectedImage := adjust.Contrast(gammaCorrectedImage, 0.2)

//		file.ConvertedImage = contrastCorrectedImage
//	}
var conversionFilters = map[string]gift.Filter{

	"invert": gift.Invert(),
	"color_balance": gift.ColorFunc(
		func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
			r = r0 + 0.17
			g = g0 + 0.02
			b = b0 - 0.06
			a = a0
			return r, g, b, a
		},
	),
	"gamma":    gift.Gamma(0.2),
	"contrast": gift.Contrast(10),
}

func (file *ImageFile) Conversion() {
	convertedImage := file.InitImage

	for name, filter := range conversionFilters {
		if name == "color_balance" && !file.IsColor {
			continue
		}

		g := gift.New(filter)
		dst := image.NewNRGBA(g.Bounds(convertedImage.Bounds()))
		g.Draw(dst, convertedImage)

		convertedImage = dst.SubImage(dst.Rect)
	}

	file.ConvertedImage = convertedImage
}
