// lissajous_1.go generates an animated GIF of Lissajous figures and writes it to standard output.
//
// Context: The Go Programming Language, Chapter 1, Animated GIFs section
// Greg Tate
// 2025-05-05

// Run go build lissajous_1.go to create an executable
// Then run .\lissajous_1.exe > lissajous.gif to create the animated GIF

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

// This block defines the color palette and constants for color indices
var palette = []color.Color{color.White, color.Black} 	// Composite literal for color.Color slice

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of frames in the animation
		delay   = 8     // delay between frames in 10ms units
	)

	freq  := rand.Float64() * 3.0           // relative frequency of y oscillator
	anim  := gif.GIF{LoopCount: nframes}	// composite literal for gif.GIF struct
	phase := 0.0                            // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi;  t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)  // Note: using gif.EncodeAll instead of gif.Encode
}