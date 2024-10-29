package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "ab c d  e   \t  f"
	a = string(squashSpaces([]byte(a)))
	fmt.Println(a)
}

func squashSpaces(s []byte) []byte {
	out := s[:0]
	for _, r := range string(s) {
		if unicode.IsSpace(r) {
			if ' ' == (rune(out[len(out)-1])) {
				continue
			} else {
				out = append(out, byte(' '))
			}
		} else {
			out = append(out, byte(r))
		}
	}
	return out
}
