package main

import (
	"log"
	"os"

	"github.com/joshuaskatz/imageproc/internal/image"
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
	path := homedir + "/desktop/lomo.tif"

	file := image.NewFile()

	file.SetFilePath(path)

	file.LoadImage()

	file.SetIsColor(true)

	file.Conversion()

	file.SaveImage()
}
