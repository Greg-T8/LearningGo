/*
Program: Echo Command-Line Arguments (Range Version)
Context: The Go Programming Language, Chapter 1
Display Name: Greg Tate
Date: April 19, 2025

This program prints its command-line arguments, separated by spaces, to standard output. This version uses a range-based for loop.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize result and separator strings
	s, sep := "", ""
	// Concatenate command-line arguments using a range-based loop
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// Output the concatenated arguments
	fmt.Println(s)
}
