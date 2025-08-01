// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	// Optional command-line flags to control the output.

	// flag.Bool() arguments: <name>, <default value>, <help message>; returns a pointer to a bool
	n = flag.Bool("n", false, "omit trailing newline")

	// flag.String() arguments: <name>, <default value>, <help message>; returns a pointer to a string
	sep = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) // Non-flag arguments are accessed via flag.Args();
	if !*n {                                   // n is a pointer to a bool, so we dereference it with *
		fmt.Println()
	}
}
