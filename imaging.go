package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func createpng(i int) {

	dc := gg.NewContext(128, 128)

	dc.DrawRectangle(0, 0, 128, 128)
	dc.SetRGB(255, 255, 255)
	dc.Fill()
	dc.SetRGB(0x00, 0x00, 0x00)
	for y := 0; y < Program.WorldY; y++ {
		for x := 0; x < Program.WorldX; x++ {
			v := World[XYtoIndex(x, y)]

			if v != 0 && v != -2 {
				dc.SetPixel(x, y)
			}
			if v == -2 {
				dc.SetPixel(x, y)
			}
		}
	}
	s := fmt.Sprintf("images/image-%08d.png", i)
	dc.SavePNG(s)

	/*
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
	*/
}
