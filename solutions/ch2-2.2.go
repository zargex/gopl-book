// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unitconv"

	"gopl.io/ch2/tempconv"
)

func main() {

	if len(os.Args) > 1 {

		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			unitTable(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			u, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			unitTable(u)
		}
	}
}

func unitTable(u float64) {
	f := tempconv.Fahrenheit(u)
	c := tempconv.Celsius(u)
	k := unitconv.Kilograms(u)
	m := unitconv.Meters(u)
	lb := unitconv.Pounds(u)
	feet := unitconv.Feets(u)
	fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c), k, unitconv.KtoP(k), m, unitconv.MtoF(m), lb, unitconv.PtoK(lb), feet, unitconv.FtoM(feet))
}

//!-
