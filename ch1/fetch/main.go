// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url,"http://") {
			fmt.Println("it don't have the prefix!!, lets fix")
			url = "http://" + url
		} else {
			fmt.Println("it do got the prefix")
		}

		resp, err := http.Get(url)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Print("Resp Type: ")
		fmt.Println(reflect.TypeOf(resp))
		fmt.Print("resp.Header Type: ")
		fmt.Println(reflect.TypeOf(resp.Header))
		fmt.Print("Status: ")
		fmt.Println(resp.Status)
		fmt.Println("Proto: ")
		fmt.Println(resp.Proto)
		fmt.Println("resp.Body type")
		fmt.Println(reflect.TypeOf(resp.Body))

		b, err := ioutil.ReadAll(resp.Body)
		//length, err := io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("type: %v and value: %s",reflect.TypeOf(length), length)
		//fmt.Printf("type: %v and value: %v",reflect.TypeOf(err), err)

		fmt.Printf("%s\n", b[:50])
		fmt.Println(reflect.TypeOf(b))
	}
}

//!-
