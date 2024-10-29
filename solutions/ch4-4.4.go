package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = rotate(a)
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int) []int {
	s = append(s[2:], s[:2]...)
	return s
}
