// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var num int
	num = 0
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		num += 1
	}
	fmt.Println(s)
	fmt.Println("The first argument: " + os.Args[1])
	fmt.Println("the number of command line arguments is: " + string(num))
	fmt.Println(num)
}

//!-
