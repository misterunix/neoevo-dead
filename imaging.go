package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func createpng2(i int) {

	dc := gg.NewContext(128, 128)

	dc.DrawRectangle(0, 0, 128, 128)
	dc.SetRGB(55, 55, 55)
	dc.Fill()

	dc.SetRGB(0x00, 0xFF, 0xFF)

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

// createpng : Creates an image from the Worldmap.
func createpng(i int) {

	dc := gg.NewContext(655, 655)

	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, 655, 655)
	dc.Fill()
	dc.SetRGB(.8, .8, .8)
	dc.DrawRectangle(0, 0, 654, 654)
	dc.SetLineWidth(5)
	dc.Stroke()
	//dc.DrawRectangle(5, 5, 645, 645)
	//dc.SetRGB(0, 0, 0)
	//dc.Fill()
	dc.SetRGB(1, 1, 1)
	for y := 0; y < Program.WorldY; y++ {
		for x := 0; x < Program.WorldX; x++ {
			v := World[XYtoIndex(x, y)]
			x11 := float64(x)*5.0 + 5.0
			y11 := float64(y)*5.0 + 5.0
			//x1 := x*5 + 5
			//y1 := y*5 + 5

			if v != 0 && v != -2 {
				dc.DrawCircle(x11, y11, 2.5)

				dc.Fill()
				//	dc.SetPixel(x1, y1)
			}
			if v == -2 {
				//dc.DrawCircle(x11, y11, 2.5)
				//dc.SetPixel(x, y)
			}
		}
	}
	s := fmt.Sprintf("images/image-%08d.png", i)
	dc.SavePNG(s)

}
