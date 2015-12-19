package main

import "fmt"

func main() {
  a := []int{0, 1, 2, 3, 4, 5}
  b := rotate(a, 2)
  fmt.Println(b)
}

func rotate(a []int, pos int) []int {
  return append(a[pos:], a[:pos]...)
}

func reverse(a *[6]int) {
  for i, j := 0, len(*a) - 1; i < j; i, j = i + 1, j - 1 {
    a[i], a[j] = a[j], a[i]
  }
}
