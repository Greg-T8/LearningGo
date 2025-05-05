// dup2 -- identifies duplicate lines in input files or stdin
//
// This program reads from standard input or from a list of named files
// and reports the count and text of lines that appear more than once.
// It uses a map to count line occurrences.
//
// From "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan
// Chapter 1 - A variant of dup1 that reads from stdin or a list of files
//
// Author: Greg Tate
// Date: 2025-04-26
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main creates a map to track line counts, processes input from stdin or files,
// and then displays any lines that appear more than once
func main() {
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

// countLines reads input lines from a file and counts occurrences in the map
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
