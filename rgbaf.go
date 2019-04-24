package main

import "image/color"

// RGBAF represents itself underneath as floating points for math purposes
type RGBAF [4]float64

// RGBA color
func (c RGBAF) RGBA() (r, g, b, a uint32) {
	return uint32(c[0]), uint32(c[1]), uint32(c[2]), uint32(c[3])
}

// BlackRGBAF Returns the color black in RGBAF
func BlackRGBAF() RGBAF {
	return [4]float64{0, 0, 0, 0}
}

// AddColor returns a new color that is the sum of the two colors
func (c RGBAF) AddColor(otherColor color.Color) RGBAF {
	other := ToRGBAF(otherColor)
	return [4]float64{
		c[0] + other[0],
		c[1] + other[1],
		c[2] + other[2],
		c[3] + other[3],
	}
}

// Divide divides each component of rgba by the divisor
func (c RGBAF) Divide(d float64) RGBAF {
	return [4]float64{
		c[0] / d,
		c[1] / d,
		c[2] / d,
		c[3] / d,
	}
}

// ToRGBAF constructs RGBAF from a color
func ToRGBAF(c color.Color) RGBAF {
	r, g, b, a := c.RGBA()
	return [4]float64{
		float64(r >> 8),
		float64(g >> 8),
		float64(b >> 8),
		float64(a >> 8),
	}
}
