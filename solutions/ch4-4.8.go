// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
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
	unicodeCategories := map[string]int{"number": 0, "space": 0, "letter": 0, "symbol": 0, "punct": 0}
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

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
		}
		counts[r]++
		utflen[n]++
		addRuneToCategory(unicodeCategories, r)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nCategory\tcount\n")
	for k, v := range unicodeCategories {
		if v > 0 {
			fmt.Printf("%s\t\t%d\n", k, v)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func addRuneToCategory(m map[string]int, r rune) {
	switch {
	case unicode.IsLetter(r):
		m["letter"]++
	case unicode.IsNumber(r):
		m["number"]++
	case unicode.IsSpace(r):
		m["space"]++
	case unicode.IsSymbol(r):
		m["symbol"]++
	case unicode.IsPunct(r):
		m["punct"]++
	}
}

//!-
