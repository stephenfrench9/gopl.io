// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
        counts := make(map[string]int)

	for _,f := range os.Args[1:] {
		file, err := os.Open(f)

		if err != nil {
		   fmt.Println("dang cousin, it broke")
		   fmt.Println(err)
		   fmt.Printf("dup2: %v\n",err)
		   fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		   continue
	   	}

		countLines(file, counts)
		fmt.Println(counts)

		for k,v := range counts {
		    if v > 1 {
           	    	    fmt.Print(k)
	    		    fmt.Print(" ")
	    	    	    fmt.Println(v)
	    	    }
        	}
	}		
}

func dmain() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
