// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for k, arg := range os.Args[1:] {
		s += sep + arg + strconv.Itoa(k)
		sep = " "
		fmt.Print(k)
	}
	fmt.Println(s)

}

//!-
