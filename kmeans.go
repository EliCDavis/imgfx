package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

func kmeans(in *image.RGBA, k, iterations int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, in.Bounds().Dx(), in.Bounds().Dy()))
	draw.Draw(m, m.Bounds(), in, image.Point{0, 0}, draw.Src)

	means := make([]color.Color, k)

	for i := 0; i < k; i++ {
		xpos := rand.Intn(in.Bounds().Dx())
		ypos := rand.Intn(in.Bounds().Dy())
		means[i] = m.At(xpos, ypos)
	}

	for timesRan := 0; timesRan < iterations; timesRan++ {

		totals := make([]RGBAF, k)
		for t := 0; t < k; t++ {
			totals[t] = BlackRGBAF()
		}

		totalsCount := make([]int, k)

		for x := 0; x < in.Bounds().Dx(); x++ {
			for y := 0; y < in.Bounds().Dy(); y++ {

				nearestDistance := math.Inf(0)
				curColor := m.At(x, y)
				nearestMeans := 0

				for meanIndex, mean := range means {
					dist := ColorDistance(curColor, mean)

					if dist < nearestDistance {
						nearestDistance = dist
						nearestMeans = meanIndex
					}
				}

				totalsCount[nearestMeans] = totalsCount[nearestMeans] + 1
				totals[nearestMeans] = totals[nearestMeans].AddColor(curColor)

			}
		}

		for meanIndex := range means {
			n := totalsCount[meanIndex]
			if n > 0 {
				means[meanIndex] = totals[meanIndex].Divide(float64(n))
			} else {
				xpos := rand.Intn(in.Bounds().Dx())
				ypos := rand.Intn(in.Bounds().Dy())
				means[meanIndex] = m.At(xpos, ypos)
			}
		}

	}

	for x := 0; x < in.Bounds().Dx(); x++ {
		for y := 0; y < in.Bounds().Dy(); y++ {
			nearestDistance := math.Inf(0)
			nearestMeans := -1
			for meanIndex, mean := range means {
				dist := ColorDistance(m.At(x, y), mean)
				if dist < nearestDistance {
					nearestDistance = dist
					nearestMeans = meanIndex
				}
			}
			m.Set(x, y, means[nearestMeans])
		}
	}

	return m
}
