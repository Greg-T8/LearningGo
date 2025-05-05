// Program: dup1.go
// Description: This program reads lines of text from standard input and prints each line that appears more than once,
//              along with the count of its occurrences. It demonstrates basic usage of maps and input scanning in Go.
// Context: Chapter 1 of "The Go Programming Language" book.
// Author: Greg Tate
// Date: 2025-04-26

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Main function: Reads input, counts duplicate lines, and prints results.
func main() {
	// Create a map to store line counts
	counts := make(map[string]int)

	// Initialize a scanner to read from standard input
	input := bufio.NewScanner(os.Stdin)

	// Loop to read each line of input
	for input.Scan() {
		counts[input.Text()]++
	}

	// Iterate over the map to find duplicates and print results
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
