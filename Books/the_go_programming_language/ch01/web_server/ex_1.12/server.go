// Server4: Combining the web server with the lissajous function

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := ""
		fmt.Println(r.FormValue("cycles"))
		cycles = r.FormValue("cycles")
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	// This block defines the color palette and constants for color indices
	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xFF},
		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		color.RGBA{0xFF, 0x00, 0x00, 0xFF},
		color.RGBA{0x00, 0xFF, 0x00, 0xFF},
		color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	}
	const (
		whiteIndex = 0
		blackIndex = 1
		redIndex   = 2
		greenIndex = 3
		blueIndex  = 4
	)
	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of frames in the animation
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // composite literal for gif.GIF struct
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: using gif.EncodeAll instead of gif.Encode
}
