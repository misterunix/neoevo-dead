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

}
