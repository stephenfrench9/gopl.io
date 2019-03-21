// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	my_main()

//	//!+array
//	a := [...]int{0, 1, 2, 3, 4, 5}
//	reverse(a[:])
//	fmt.Println(a) // "[5 4 3 2 1 0]"
//	//!-array
//
//	//!+slice
//	s := []int{0, 1, 2, 3, 4, 5}
//	// Rotate s left by two positions.
//	reverse(s[:3])
//	reverse(s[3:])
//	reverse(s)
//	fmt.Println(s) // "[2 3 4 5 0 1]"
//	//!-slice
//
//	// Interactive test of reverse.
//	input := bufio.NewScanner(os.Stdin)
//outer:
//	for input.Scan() {
//		var ints []int
//		for _, s := range strings.Fields(input.Text()) {
//			x, err := strconv.ParseInt(s, 10, 64)
//			if err != nil {
//				fmt.Fprintln(os.Stderr, err)
//				continue outer
//			}
//			ints = append(ints, int(x))
//		}
//		reverse(ints)
//		fmt.Printf("%v\n", ints)
//	}
//	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for    i, j := 0, len(s)-1      ;   i < j  ;    i, j = i+1, j-1    {
		s[i], s[j] = s[j], s[i]
	}
}
//!-rev

//can I take a slice, manipulate it, and then see the underlying array be manipulated as well? yes you can
//What if I replace some elements in the slice? What if I add to the end?

func my_main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := strings.Fields(input.Text())
		fmt.Println(s[1])
	}
//outer:
//	for input.Scan() {
//		var ints []int
//		for _, s := range strings.Fields(input.Text()) {
//			x, err := strconv.ParseInt(s, 10, 64)
//			if err != nil {
//				fmt.Fprintln(os.Stderr, err)
//				continue outer
//			}
//			ints = append(ints, int(x))
//		}
//		reverse(ints)
//		fmt.Printf("%v\n", ints)
//	}
//
//
}