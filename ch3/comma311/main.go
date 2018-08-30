// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaWithDot(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var c byte
	n := len(s)
	if n >= 2 {
		switch {
		case strings.HasPrefix(s, "-"):
			c = '-'
		case strings.HasPrefix(s, "+"):
			c = '+'
		}
	}
	if n <= 3 || n <= 4 && c != 0 {
		return s
	}
	if c != 0 {
		return string(c) + comma(s[1:n-3]) + "," + s[n-3:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithDot(s string) string {
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	} else {
		return comma(s)
	}
}

//!-
