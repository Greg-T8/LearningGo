// Program: Fetch URL Content
// Context: The Go Programming Language, Chapter 1, Exercise 1.8
// Author: Greg Tate
// Date: 2024-05-14

// Objective: Modify fetch to add the prefix http:// to each argument URL if it
// is missing. You might want to use strings.HasPrefix

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Loop through each URL provided as a command-line argument
	for _, url := range os.Args[1:] {
		// Prepend "https://" to the URL if it does not already have it
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		// Send an HTTP GET request to the URL and handle errors
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Copy the response body to standard output
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}