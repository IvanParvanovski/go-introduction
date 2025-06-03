package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// var palette = color.RGBA{0x00, 0xFF, 0x00, 0x00}
var palette = []color.Color{
	color.Black,                                     // black
	color.White,                                     // white
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},             // green
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},             // red
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},             // blue
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},             // yellow
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},             // cyan
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},             // magenta
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF},             // orange
	color.RGBA{0x80, 0x00, 0x80, 0xFF},             // purple
}

const (
	min = 2 
	max = 9
)

const (
	blackIndex   = 0
	whiteIndex   = 1
	greenIndex   = 2
	redIndex     = 3
	blueIndex    = 4
	yellowIndex  = 5
	cyanIndex    = 6
	magentaIndex = 7
	orangeIndex  = 8
	purpleIndex  = 9
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001 
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			colorNumber := uint8(rand.Intn(max - min) + min)
			img.SetColorIndex(size + int(x * size + 0.5), size+int(y*size+0.5), colorNumber)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
