package image

import (
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"
)

func (i *ImageFile) Conversion() {
	err := i.vipsRef.Invert()

	if err != nil {
		fmt.Println(err)
	}

	i.vipsRef.TransformICCProfile(vips.SRGBIEC6196621ICCProfilePath)
	if err != nil {
		fmt.Println(err)
	}

	exportParams := vips.NewDefaultExportParams()

	convertedImage, err := i.vipsRef.ToImage(exportParams)

	if err != nil {
		fmt.Println(err)
	}
	i.ConvertedImage = convertedImage
}
