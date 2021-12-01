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
	g := color.NRGBA{
		R: uint8(0),
		G: uint8(255),
		B: uint8(0),
		A: 255,
	}
	n := 0
	food := 0
	z := 0
	for y := 0; y < Program.WorldY; y++ {
		for x := 0; x < Program.WorldX; x++ {
			v := World[XYtoIndex(x, y)]
			z++
			if v != 0 && v != -2 {
				img.Set(x, y, c)
				n++
			}
			if v == -2 {
				img.Set(x, y, g)
				food++
			}

		}
	}

	fmt.Println(n, food, z)

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
