package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func average(colors [][]color.Color) color.Color {
	rTot, gTot, bTot, aTot := uint32(0), uint32(0), uint32(0), uint32(0)
	count := uint32(0)
	for _, column := range colors {
		for _, color := range column {
			r, g, b, a := color.RGBA()
			rTot += r >> 8
			gTot += g >> 8
			bTot += b >> 8
			aTot += a >> 8
			count++
		}
	}
	return color.RGBA{
		R: uint8((rTot / count)),
		G: uint8((gTot / count)),
		B: uint8((bTot / count)),
		A: uint8(255),
	}
}

// TODO: kernel doesn't do proper wrapping for size greater than 1
func kernel(xCenter int, yCenter int, size int, context *image.RGBA) [][]color.Color {
	kernel := make([][]color.Color, (size*2)+1)

	yIter := 0
	xIter := 0

	for y := yCenter - size; y <= yCenter+size; y++ {
		kernel[yIter] = make([]color.Color, (size*2)+1)
		xIter = 0
		yCord := y
		if y < 0 {
			yCord = yCenter + size
		}
		if y >= context.Bounds().Dy() {
			yCord = yCenter - size
		}
		for x := xCenter - size; x <= xCenter+size; x++ {
			xCord := x
			if x < 0 {
				xCord = xCenter + size
			}
			if x >= context.Bounds().Dx() {
				xCord = xCenter - size
			}
			kernel[yIter][xIter] = context.At(xCord, yCord)
			xIter++
		}

		yIter++
	}
	return kernel
}

func smooth(in *image.RGBA, amount int) *image.RGBA {
	out := image.NewRGBA(image.Rect(0, 0, in.Bounds().Dx(), in.Bounds().Dy()))
	draw.Draw(out, out.Bounds(), in, image.Point{0, 0}, draw.Src)

	for y := 0; y < in.Bounds().Dy(); y++ {
		for x := 0; x < in.Bounds().Dx(); x++ {
			out.Set(x, y, average(kernel(x, y, amount, in)))
		}
	}

	return out
}

// func wavify(context *gg.Context, offset int) *gg.Context {

// 	blurred := smooth(context, 3)
// 	out := gg.NewContext(context.Width(), context.Height())

// 	maxBrightness := 300000
// 	for y := 0; y < out.Height(); y++ {

// 		brightnessCount := float64(offset)

// 		for x := 0; x < out.Width(); x++ {
// 			brightnessCount += brightness(blurred.Image().At(x, y))
// 			if brightnessCount > float64(maxBrightness) {
// 				out.SetRGB(1, 1, 1)
// 				brightnessCount -= float64(maxBrightness / 2)
// 			} else {
// 				out.SetRGB(0, 0, 0)
// 			}
// 			out.SetPixel(x, y)
// 		}
// 	}

// 	return out
// }

func main() {

	filename := "resistance.png"

	imgfile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer imgfile.Close()

	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err)
	}

	toimg, _ := os.Create("out.png")
	defer toimg.Close()

	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(m, m.Bounds(), img, image.Point{0, 0}, draw.Src)

	png.Encode(toimg, kmeans(smooth(m, 3), 20, 200))

}
