// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Printf("You need just two words to compare.")
		os.Exit(-1)
	}
	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s and %s are anagrams!", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s and %s are not anagrams", os.Args[1], os.Args[2])
	}
}

func isAnagram(s1, s2 string) bool {
	if utf8.RuneCountInString(s2) != utf8.RuneCountInString(s1) {
		return false
	}
	table1 := make(map[rune]int)
	table2 := make(map[rune]int)
	for _, r := range s1 {
		table1[r] = table1[r] + 1
	}
	for _, r := range s2 {
		table2[r] = table2[r] + 1
	}
	for _, r := range s1 {
		if table1[r] != table2[r] {
			return false
		}
	}

	return true
}
