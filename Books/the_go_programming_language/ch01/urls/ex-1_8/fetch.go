// fetch.go fetches each URL provided as a command-line argument, prepending "http://" if missing, and streams the response to standard output.
//
// Context: The Go Programming Language, Exercise 1.8
// Greg Tate
// 2025-05-14

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// main is the entry point of the program.
// It iterates over command-line arguments (URLs), ensures each has an "https://" prefix,
// fetches the content using HTTP GET, and writes the response body to standard output.
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}