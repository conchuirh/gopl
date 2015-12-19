package main

import "fmt"

func main() {
  s := []string{"hello", "hello", "hello", "yes", "yes", "no"}
  fmt.Println(removeDup(s))
}

func removeDup(s []string) []string {
  prev := s[0]
  for i := 1; i < len(s); i++ {
    if prev == s[i] {
      prev = s[i]
      if i + 1 < len(s) {
        s = append(s[:i], s[i + 1:]...)
        i--
      }
    } else {
      prev = s[i]
    }
  }
  return s
}
