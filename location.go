package main

import "math"

// GetDistance : Return the distance between x1,y1 to x2,y2
func GetDistance(x1, y1, x2, y2 float64) float64 {
	d := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
	return d
}

// GetAngle : Return the angle if 0.0 to 1.0 from x1,y1 to x2,y2
func GetAngle(x1, y1, x2, y2 float64) float64 {
	a := math.Atan2(y2-y1, x2-x1) * RAD2DEG
	if a < 0 {
		a += 360.0
	}
	if a >= 360.0 {
		a -= 360.0
	}

	return a / 360.0
}

// Convert X,Y to index value for positioning in the world space.
func XYtoIndex(x, y int) int {

	// sanity checks
	if x >= Program.WorldX {
		x = Program.WorldX - 1
	}
	if y >= Program.WorldY {
		y = Program.WorldY - 1
	}
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	return y*Program.WorldX + x
}

// IndexToXY : Convert the index into X,Y for the world space.
func IndexToXY(i int) Point {
	// sanity checks
	if i >= Program.WorldSize {
		i = Program.WorldSize - 1
	}

	p := Point{
		Y: i / Program.WorldX,
		X: i % Program.WorldX,
	}
	return p
}
