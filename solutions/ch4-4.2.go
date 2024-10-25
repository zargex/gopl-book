// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var alg = flag.String("algo", "sha256", "which algorith to use (sha256, sha384 or sha512)")

func main() {
	flag.Parse()
	if *alg != "sha256" && *alg != "sha384" && *alg != "sha512" {
		fmt.Println("Wrong flag, use '-h' for help")
		os.Exit(-1)
	}

	if *alg == "sha256" {
		for _, arg := range flag.Args() {
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		}
	} else if *alg == "sha384" {
		for _, arg := range flag.Args() {
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		}
	} else {
		for _, arg := range flag.Args() {
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		}
	}

}

//!-
