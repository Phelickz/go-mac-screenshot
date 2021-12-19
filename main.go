package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("starting screenshot")

	displays := screenshot.NumActiveDisplays()

	for i := 0; i < displays; i++ {
		var images string
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			panic(err)
		}

		fmt.Print("What will you name it? : ")
		fmt.Scanln(&images)
		fileName := images
		fileName = fileName + ".png"

		file, _ := os.Create(fileName)

		defer file.Close()

		png.Encode(file, img)

		fmt.Printf("%s\n", fileName+".png")
	}
}
