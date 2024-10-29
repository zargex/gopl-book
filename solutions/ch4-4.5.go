package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c", "c", "d", "e", "e", "f"}
	a = removeDuplicate(a)
	fmt.Println(a)
}

func removeDuplicate(s []string) []string {
	out := s[:0]
	for i := 0; i < len(s); i++ {
		j := i + 1
		if j == len(s) {
			out = append(out, s[i])
		} else {
			if s[i] == s[j] {
				out = append(out, s[i])
				i++
			} else {
				out = append(out, s[i])
			}
		}
	}
	return out
}

func rotate(s []int) []int {
	s = append(s[2:], s[:2]...)
	return s
}
