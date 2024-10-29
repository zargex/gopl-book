package main

import (
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', '\xe4', '\xb8', '\x96'}
	a = reverse(a)
	fmt.Printf("%s\n", a)
}

func reverse(s []byte) []byte {
	runes := []rune(string(s))

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
