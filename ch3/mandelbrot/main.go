package main

import (
  "image"
  "image/color"
  "image/png"
  "math/cmplx"
  "os"
)

var total complex128

func main() {
  const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2
    width, height          = 1024, 1024
  )

  img := image.NewRGBA(image.Rect(0, 0, width, height))
  var totalR float64
  var totalG float64
  for py := 0; py < height; py++ {
    for px := 0; px < width; px++ {
      for i := 0.0; i <= 1.0; i += .25 {
        y := (float64(py) + i) / height * (ymax - ymin) + ymin
        x := (float64(px) + i) / width * (xmax - xmin) + xmin
        z := complex(x, y)
        r, g := mandelbrot(z)
        totalR += r
        totalG += g
      }
      totalR = (totalR) / 4
      totalG = (totalG) / 4
      img.Set(px, py, color.RGBA{uint8(totalR), uint8(totalG), 0, 0xff})
    }
  }
  png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) (float64,float64) {
  const iterations = 200
  const contrast = 15
  total++;
  var v complex128
  for n := uint8(0); n < iterations; n++ {
    v = v * v + z
    if cmplx.Abs(v) > 2 {
      //ex5
      return float64(255 - contrast * n), float64(0 + contrast * n)
    }
  }
  return 0, 0
}
