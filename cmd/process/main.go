package main

import (
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
	path := homedir + "/desktop/decay.tif"

	file := internal_image.NewFile()

	file.SetFilePath(path)

	file.LoadImage()

	file.SetIsColor(false)

	file.Conversion()

	file.SaveImage()
}
