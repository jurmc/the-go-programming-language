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

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 1*32 - 1, A: 255},
	color.RGBA{R: 2*32 - 1, A: 255},
	color.RGBA{R: 3*32 - 1, A: 255},
	color.RGBA{R: 4*32 - 1, A: 255},
	color.RGBA{R: 5*32 - 1, A: 255},
	color.RGBA{R: 6*32 - 1, A: 255},
	color.RGBA{R: 7*32 - 1, A: 255},
	color.RGBA{R: 8*32 - 1, A: 255},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillatro revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image cnavas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			dist := math.Sqrt(x*x + y*y)
			col_idx := uint8(dist * 8)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				col_idx)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
