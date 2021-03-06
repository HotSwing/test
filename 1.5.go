package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"net/http"
	"time"
	"log"
)

var palette = []color.Color{color.Black, color.RGBA{0xFF, 0xFF, 0x0, 0xFF},color.RGBA{0x0, 0xFF, 0xFF, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, R *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
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
	flag := 0
	for i := 0; i < nframes ; i++ {
		rect := image.Rect(0, 0, 2*size+1, 3*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0 ; t < cycles*2*math.Pi ; t += res {
			x := math.Cos(t)
			y := math.Sin(t*freq + phase)
			if flag % 2 == 0 {
			img.SetColorIndex(size+int(x*size+0.5), size +int(y*size+0.5),
			1)
			} else {
			img.SetColorIndex(size+int(x*size+0.5), size +int(y*size+0.5),
			2)	
			}
			flag ++
		}
		phase +=0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}