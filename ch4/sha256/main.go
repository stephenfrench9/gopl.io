// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

//Stephen: You can format outputs by using fmt.Printf, and then you put a number after the percent sign.
//Stephen: A uint8 can be represented as a hexint, or as a string.
//Stephen: you can have an array of unint8s.
//Stephen: is == superficial? It seems like yes it is, it checks values, not the underlying object.


import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//c1 := "x"
	//c2 := "X"
	//c1 := []byte("x")
	//c2 := []byte("X")
	//c1 := []byte("x")[1]
	//c2 := []byte("X")[1]
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("string: %4s\nstring: %4s\nhex int: %4x\nhex int: %4x\nSame: %7t\ntype: %7T\n", c1, c2, c1, c2, c1 == c2, c1)


	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-

