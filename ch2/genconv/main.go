package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/conchuirh/gopl/ch2/tempconv"
  "github.com/conchuirh/gopl/ch2/conv"
)

func main() {
  if len(os.Args) == 1 {
    var input string
    fmt.Scanln(&input)
    t, err := strconv.Atoi(input)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }
    f := tempconv.Fahrenheit(t)
    c := tempconv.Celsius(t)
    feet := conv.Feet(t)
    meters := conv.Meters(t)
    pounds := conv.Pounds(t)
    kg := conv.Kilograms(t)
    fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
    fmt.Printf("%s = %s, %s = %s\n", feet, conv.FToM(feet), meters, conv.MToF(meters))
    fmt.Printf("%s = %s, %s = %s\n", pounds, conv.PToK(pounds), kg, conv.KToP(kg))
  } else {
    for _, arg := range os.Args[1:] {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        os.Exit(1)
      }
      f := tempconv.Fahrenheit(t)
      c := tempconv.Celsius(t)
      feet := conv.Feet(t)
      meters := conv.Meters(t)
      pounds := conv.Pounds(t)
      kg := conv.Kilograms(t)
      fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
      fmt.Printf("%s = %s, %s = %s\n", feet, conv.FToM(feet), meters, conv.MToF(meters))
      fmt.Printf("%s = %s, %s = %s\n", pounds, conv.PToK(pounds), kg, conv.KToP(kg))
    }
  }
}
