package main

import (
	"image/color"
	"math"
)

// ColorDistance calculates the distance between two colors by treating the
// rgb values as points in space. (Ignores alpha)
func ColorDistance(c1 color.Color, c2 color.Color) float64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	r := math.Pow(float64(r1)-float64(r2), 2.0)
	g := math.Pow(float64(g1)-float64(g2), 2.0)
	b := math.Pow(float64(b1)-float64(b2), 2.0)

	return math.Sqrt(r + g + b)
}

// Brightness determines an arbitrary value meant to represent how bright a
// pixel is. The higher the better
func Brightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	return math.Sqrt(math.Pow(0.299*float64(r), 2.0) + math.Pow(0.587*float64(g), 2.0) + math.Pow(0.114*float64(b), 2.0))
}
