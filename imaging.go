package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

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

	dc.SetRGB(1, 1, 1)
	for y := 0; y < Program.WorldY; y++ {
		for x := 0; x < Program.WorldX; x++ {
			v := World[XYtoIndex(x, y)]
			x11 := float64(x)*5.0 + 5.0
			y11 := float64(y)*5.0 + 5.0

			if v != 0 && v != -2 {
				dc.DrawCircle(x11, y11, 2.5)
				dc.Fill()
			}
			//if v == -2 {
			//	}
		}
	}
	s := fmt.Sprintf("images/gen-%06d-%08d.png", generation, i)
	dc.SavePNG(s)

}
