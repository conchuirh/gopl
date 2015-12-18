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
  "strings"
  "strconv"
)

func main() {
  http.HandleFunc("/", handler) // each request call handler
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  var (
    cycles  = 5
    size    = 100
    nframes = 64
    delay   = 8
  )
  list := strings.Split(r.URL.String(), "=")
  switch list[0] {
  case "?cycles":
    cycles, _ = strconv.Atoi(list[1])
  case "?size":
    size, _ = strconv.Atoi(list[1])
  case "?nframes":
    nframes, _ = strconv.Atoi(list[1])
  case "delay":
    delay, _ = strconv.Atoi(list[1])
  }
  lissajous(w, cycles, size, nframes, delay)
}

var palette = []color.Color{color.White, color.Black}

const (
  whiteIndex = 0
  blackIndex = 1
)

func lissajous(out io.Writer, c, s, nf, d int) {
  var (
    cycles  = c
    res     = 0.001
    size    = s
    nframes = nf
    delay   = d
  )
  freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
