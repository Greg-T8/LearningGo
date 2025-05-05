// dup4 reports counts and filenames of duplicate lines from input files.
// It reads from stdin or from a list of named files and prints the count and
// filenames for each line that appears more than once.
//
// Context: From "The Go Programming Language" book, Chapter 1, Section 1.4
// This is extension of dup2 example that also tracks filenames.
//
// Greg Tate
// 2025-04-26
package main

import (
	"bufio"
	"fmt"
	"os"
)

// The main function initializes data structures and processes files
func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

// countLines reads input lines from a file and counts occurrences in the map
func countLines(f *os.File, counts map[string]int, filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filenames[input.Text()] += f.Name() + ", "
	}
	// NOTE: ignoring potential errors from input.Err()
}
