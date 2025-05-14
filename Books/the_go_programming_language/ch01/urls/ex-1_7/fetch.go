// fetch.go streams the content of each URL provided as a command-line argument directly to standard output using io.Copy.
//
// Context: The Go Programming Language, Exercise 1.7
// Greg Tate
// 2025-05-14

// Objective: The function call `io.copy(dst, src)` reads from src and writes to
// dst. use it instead of `ioutil.readall` to copy the response body to
// os.stdout without requiring a buffer large enough to hold the entire stream.
// be sure to check the error result of io.copy.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// Make HTTP GET request to the URL and store result in resp struct
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Read the response and write it to stdout using io.Copy
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}