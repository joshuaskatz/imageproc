package internal_image

import (
	"image"

	"github.com/disintegration/gift"
)

var conversionFilters = map[string]gift.Filter{
	"color_balance": gift.ColorFunc(
		func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
			r = r0 + 0.15
			g = g0
			b = b0 - 0.05
			a = a0
			return r, g, b, a
		},
	),
	"invert":   gift.Invert(),
	"contrast": gift.Contrast(10),
	"gamma":    gift.Gamma(0.2),
}

func (file *File) Conversion() {
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

	file.InvertedImage = convertedImage
}
