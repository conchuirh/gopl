package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	countLetter := make(map[rune]int)
  countDigit := make(map[rune]int)
  countControl := make(map[rune]int)
  countGraphic := make(map[rune]int)
  countMark := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		} else if unicode.IsLetter(r) {
      countLetter[r]++
      utflen[n]++
    } else if unicode.IsDigit(r) {
      countDigit[r]++
      utflen[n]++
    } else if unicode.IsControl(r) {
      countControl[r]++
      utflen[n]++
    } else if unicode.IsGraphic(r) {
      countGraphic[r]++
      utflen[n]++
    } else if unicode.IsMark(r) {
      countMark[r]++
      utflen[n]++
    } else {
      invalid++
			continue
    }
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range countLetter {
		fmt.Printf("%q\t%d\n", c, n)
	}
  fmt.Printf("digit\tcount\n")
	for c, n := range countDigit {
		fmt.Printf("%q\t%d\n", c, n)
	}
  fmt.Printf("control\tcount\n")
	for c, n := range countControl {
		fmt.Printf("%q\t%d\n", c, n)
	}
  fmt.Printf("graphic\tcount\n")
	for c, n := range countGraphic {
		fmt.Printf("%q\t%d\n", c, n)
	}
  fmt.Printf("mark\tcount\n")
	for c, n := range countMark {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
