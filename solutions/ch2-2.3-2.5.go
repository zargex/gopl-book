package main

import "fmt"

func main() {
	fmt.Println(PopCount2(549846541324))
	//fmt.Println(PopCount3(6))
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var acc int
	for i := 0; i < 8; i++ {
		acc += int(pc[byte(x>>(i*8))])
	}
	return acc
	/*return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))])
	*/
}

func PopCount2(x uint64) int {
	var acc int
	for i := 0; i < 64; i++ {
		acc += int((x >> i) & 1)
	}
	return acc
}

func PopCount3(x uint64) int {
	var acc int
	for x > 0 {
		acc++
		x = x & (x - 1)
	}
	return acc
}
