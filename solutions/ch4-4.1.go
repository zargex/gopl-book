// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("bits differents between 'x' and 'X': %d\n", different_bits(c1, c2))
}

//!-

func different_bits(a, b [32]byte) int {
	diff := 0
	for i, _ := range a {
		diff += PopCount3(a[i] ^ b[i])
	}
	return diff
}

func PopCount3(x byte) int {
	var acc int
	for x > 0 {
		acc++
		x = x & (x - 1)
	}
	return acc
}
