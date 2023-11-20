package image

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2/data/binding"
	"github.com/disintegration/gift"
)

func (i *ImageFile) RGBChannels(rBinding, gBinding, bBinding binding.Float, refreshImage func()) {
	var g *gift.GIFT
	rListener := binding.NewDataListener(func() {
		r, err := rBinding.Get()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r)

		r0 := float32(r)

		g = gift.New(gift.ColorFunc(func(r1, g1, b1, a1 float32) (r, g, b, a float32) {

			return r0 / 100, g1, b1, a1
		}))
		dst := image.NewNRGBA(g.Bounds(i.ConvertedImage.Bounds()))
		g.Draw(dst, i.ConvertedImage)

		convertedImage := dst.SubImage(dst.Rect)

		i.ConvertedImage = convertedImage
		refreshImage()
	})
	gListener := binding.NewDataListener(func() {
		_g, err := gBinding.Get()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(_g)

		g0 := float32(_g)
		g = gift.New(gift.ColorFunc(func(r1, g1, b1, a1 float32) (r, g, b, a float32) {

			return r1, g0 / 100, b1, a1
		}))
		dst := image.NewNRGBA(g.Bounds(i.ConvertedImage.Bounds()))
		g.Draw(dst, i.ConvertedImage)

		convertedImage := dst.SubImage(dst.Rect)

		i.ConvertedImage = convertedImage
		refreshImage()
	})
	bListener := binding.NewDataListener(func() {
		b, err := bBinding.Get()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(b)

		b0 := float32(b)
		g = gift.New(gift.ColorFunc(func(r1, g1, b1, a1 float32) (r, g, b, a float32) {

			return r1, g1, b0 / 100, a1
		}))
		dst := image.NewNRGBA(g.Bounds(i.ConvertedImage.Bounds()))
		g.Draw(dst, i.ConvertedImage)

		convertedImage := dst.SubImage(dst.Rect)

		i.ConvertedImage = convertedImage
		refreshImage()

	})

	rBinding.AddListener(rListener)
	gBinding.AddListener(gListener)
	bBinding.AddListener(bListener)

}
