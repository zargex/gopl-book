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
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var decimalString string
	floatPoint := strings.LastIndex(s, ".")
	if floatPoint != -1 {
		for i := len(s) - 1; i >= floatPoint; i-- {
			buf.WriteByte(s[i])
		}
		decimalString = s[:floatPoint]
	} else {
		decimalString = s
	}

	sep := 0
	for i := len(decimalString) - 1; i >= 0; i-- {
		if sep == 3 {
			if s[i] != '+' && s[i] != '-' {
				buf.WriteString(",")
			}
			sep = 0
		}
		buf.WriteByte(decimalString[i])
		sep++
	}
	slices.Reverse(buf.Bytes())
	return buf.String()
}

//!-
