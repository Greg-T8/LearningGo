// Server4: Combining the web server with the lissaous function

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
	"strconv"
)

func main() {
	fmt.Println("Web server listening on 'localhost:8000'")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		qp_cycles := r.FormValue("cycles")
		if qp_cycles == "" {
			qp_cycles = "5" // Default value if not provided
		}
		cycles, err := strconv.Atoi(qp_cycles)
		if err != nil {
			log.Fatal(err)
		}

		qp_res := r.FormValue("res")
		if qp_res == "" {
			qp_res = "0.001" // Default value if not provided
		}
		res, err := strconv.ParseFloat(qp_res, 64)
		if err != nil {
			log.Fatal(err)
		}

		qp_size := r.FormValue("size")
		if qp_size == "" {
			qp_size = "100" // Default value if not provided
		}
		size, err := strconv.Atoi(qp_size)
		if err != nil {
			log.Fatal(err)
		}

		qp_nframes := r.FormValue("nframes")
		if qp_nframes == "" {
			qp_nframes = "64" // Default value if not provided
		}
		nframes, err := strconv.Atoi(qp_nframes)
		if err != nil {
			log.Fatal(err)
		}

		qp_delay := r.FormValue("delay")
		if qp_delay == "" {
			qp_delay = "8" // Default value if not provided
		}
		delay, err := strconv.Atoi(qp_delay)
		if err != nil {
			log.Fatal(err)
		}

		lissajous(w, cycles, res, size, nframes, delay)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
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
	// res     = 0.001 // angular resolution
	// size    = 100 // image canvas covers [-size..+size]
	// nframes = 64 // number of frames in the animation
	// delay   = 8  // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // composite literal for gif.GIF struct
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res { // Note: Go does not allow automatic type conversions between int and floats, so you must do explicit conversions
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: using gif.EncodeAll instead of gif.Encode
}
