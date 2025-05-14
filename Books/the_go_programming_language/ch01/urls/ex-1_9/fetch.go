// Program: Fetch URL Content and Display HTTP Status
// Context: The Go Programming Language, Chapter 1, Exercise 1.9
// Author: Greg Tate
// Date: 2024-06-09
// Location: Austin, TX

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
		// Print the HTTP response status code and copy the response body to standard output
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}