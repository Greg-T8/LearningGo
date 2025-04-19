/*
Program: Echo Command-Line Arguments
Context: The Go Programming Language, Chapter 1
Display Name: Greg Tate
Date: April 19, 2025

This program prints its command-line arguments, separated by spaces, to standard output.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize empty string variables for the result and separator
	var s, sep string
	// Iterate over command-line arguments (excluding the program name)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// Print the concatenated arguments
	fmt.Println(s)
}
