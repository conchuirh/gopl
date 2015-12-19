package main

import (
  "fmt"
  "crypto/sha256"
)

func main() {
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("X"))

  count := 0

  for i, arg := range c1 {
    if arg != c2[i] {
      count++
    }
  }

  fmt.Println(count)
}
