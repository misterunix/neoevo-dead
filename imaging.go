package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func createpng(i int) {

	const width, height = 128, 128

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	c := color.NRGBA{
		R: uint8(255),
		G: uint8(255),
		B: uint8(255),
		A: 255,
	}
	for y := 0; y < Program.WorldY; y++ {
		for x := 0; x < Program.WorldX; x++ {

			if World[XYtoIndex(x, y)] != 0 {
				img.Set(x, y, c)
			}
		}
	}

	s := fmt.Sprintf("images/image-%08d.png", i)
	f, err := os.Create(s)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
