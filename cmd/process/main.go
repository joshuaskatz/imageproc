package main

import (
	"fmt"
	"log"
	"os"

	internal_image "github.com/joshuaskatz/imageproc/internal/image"
)

type application struct {
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{logger: logger}

	homedir, err := os.UserHomeDir()

	if err != nil {
		app.logger.Fatal(err)
	}
	path := homedir + "/desktop/fdr_park.tif"

	file := internal_image.NewFile(path)

	fmt.Println(file.Path)

	file.LoadImage()

	invertedImage := Process(file)

	file.InvertedImage = *invertedImage

	file.SaveImage()

}
