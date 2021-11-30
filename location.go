package main

import (
	"math"
)

const (
	DEG2RAD = 0.0174532925
	RAD2DEG = 57.2957795130
)

// GetDistance : Return the distance between x1,y1 to x2,y2
func GetDistance(x1, y1, x2, y2 float64) float64 {
	d := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
	return d
}

// GetAngle : Return the angle in degrees from x1,y1 to x2,y2
func GetAngle(x1, y1, x2, y2 float64) float64 {
	a := math.Atan2(y2-y1, x2-x1) * RAD2DEG
	if a < 0 {
		a += 360
	}
	if a >= 360 {
		a -= 360
	}
	return a
}

// Convert X,Y to index value for positioning in the world space
func XYtoIndex(x, y int) int {
	return y*Program.WorldX + x
}

func IndexToXY(i int) Point {
	p := Point{
		Y: i / Program.WorldX,
		X: i % Program.WorldX,
	}
	return p
}

// DirectionToStep : Convert direction to x,y stepping
func DirectionToStep(d int) Point {
	var stepx int
	var stepy int

	switch d {
	case 0:
		stepx = 1
		stepy = 0
	case 45:
		stepx = 1
		stepy = 1
	case 90:
		stepx = 0
		stepy = 1
	case 135:
		stepx = -1
		stepy = 1
	case 180:
		stepx = -1
		stepy = 0
	case 225:
		stepx = -1
		stepy = -1
	case 270:
		stepx = 0
		stepy = -1
	case 315:
		stepx = 1
		stepy = -1
	}
	p := Point{
		X: stepx,
		Y: stepy,
	}
	return p
}
