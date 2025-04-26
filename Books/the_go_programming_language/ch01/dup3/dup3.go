// Program: dup3
// Description: Reports duplicate lines in named files by reading entire files at once
// Book Context: The Go Programming Language, Chapter 1, Section 1.3
// Author: Greg Tate
// Date: 2025-04-26

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Create a map to store line counts
	counts := make(map[string]int)

	// Process each file specified on the command line
	for _, filename := range os.Args[1:] {
		// Read entire file contents at once
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// Count occurrences of each line
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		// Display lines that appear more than once
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
