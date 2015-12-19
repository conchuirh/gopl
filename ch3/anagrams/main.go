package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(s1, s2 string) bool {
  found := true
  for i := 0; i < len(s1); i++ {
    if strings.Count(s1, string(s1[i])) != strings.Count(s2, string(s1[i])) {
      found = false
    }
  }
  return found
}
