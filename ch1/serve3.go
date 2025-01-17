package main

import (
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w,r)
	})
	http.ListenAndServe(":8000", nil)
}

var palette = []color.Color{color.White,
	color.RGBA{0x58, 0xEF, 0x7D, 0xFF},
	color.RGBA{0x94,0x00,0x00, 0xFF},
	color.RGBA{0x44, 0x4B, 0xB5, 0xFF}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	cycles := 5

	for k, v := range r.Form {
		if k=="cycles" {
			cycles,_ = strconv.Atoi(v[0])
		}
	}

	

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 练习题添加多个颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%3+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
