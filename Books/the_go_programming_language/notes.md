# My Notes from "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan

<img src='images/20250419135706.png' width='350'/>

<details>
<summary>Book Resources</summary>

- [Official Book Website](https://www.gopl.io/)
- [Source Code Examples](https://github.com/adonovan/gopl.io)

</details>

<!-- omit in toc -->
## Go Commands 

```Go
go run hello.go         // Run the program hello.go
go build hello.go       // Build the program hello.go into an executable file
```
<!-- omit in toc -->
## Contents

- [Overview and History of Go](#overview-and-history-of-go)
  - [The Go Project](#the-go-project)
- [1. Tutorial](#1-tutorial)
  - [1.2 Command-Line Arguments](#12-command-line-arguments)
  - [1.3 Funding Duplicate Lines](#13-funding-duplicate-lines)
  - [1.4 Animated GIFs](#14-animated-gifs)
  - [1.5 Fetching a URL](#15-fetching-a-url)
  - [1.6 Fetching URLs Concurrently](#16-fetching-urls-concurrently)
  - [1.7 A Web Server](#17-a-web-server)


## Overview and History of Go

**Conception**: Go was conceived in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google and was publicly announced in 2009.

**Influencers of Go**:
  
  <img src='images/20250419140256.png' width='400'/>

  - **C**: Syntax, control structures, basic data types, call-by-value parameter passing, pointers, and emphasis on compiling to machine code.
  - **Modula-2**: Inspired the package concept.
  - **Oberon**: Eliminated the distinction between module interface files and module implementation files.
  - **Oberon-2**: Influenced syntax for packages and declarations, particularly method declarations.
  - **CSP (Communicating Sequential Processes)**: Influenced goroutines and channels.
  - **Squeak**: Provided handling for mouse and keyboard events, with statically created channels.
  - **Newsqueak**: A purely functional language with garbage collection, aimed at managing keyboard, mouse, and window events.
  - **Alef**: Attempted to make Newsqueak a viable system programming language, but its lack of garbage collection made concurrency too painful.

### The Go Project

- **Motivation**: The Go project was born out of frustration with several software systems at Google that were suffering from an explosion of complexity.
- **Components**:
  - The language itself.
  - Its tools and standard libraries.
  - A cultural agenda of radical simplicity.
- **Features**:
  - Garbage collection.
  - A package system.
  - First-class functions.
  - Lexical scope.
  - A system call interface.
  - Immutable strings.
- **Minimalism**:
  - Go has comparatively few features and is unlikely to add more.
  - **Notable omissions**:
    - No implicit numeric conversions.
    - No constructors or destructors.
    - No operator overloading.
    - No default parameter values.
    - No inheritance.
    - No generics.
    - No exceptions.
    - No macros.
    - No function annotations.
    - No thread-local storage.

## 1. Tutorial

The following program can be compiled and ran with `go run hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

The command `go build hello.go` creates an executable file.

- Go code is organized into packages, which are similar to libraries or modules in other languages.
- Each source file begins with a package declaration, which states the package the file belongs go.

```Go
package main
```

- You must import *exactly* the packages you need. The Go compiler will not compile a package if it is not used in the code.
- Go does not require semicolons; they are automatically inserted at the end of lines.

### 1.2 Command-Line Arguments

The `os.Args` variable is a slice of strings containing the command-line arguments.
  - `os.Args[0]` is the name of the command.
  - `os.Args[1:]` contains the arguments passed to the program.

**Example:** Concatenating and printing command-line arguments (index-based loop):

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}
```

- `var` declares variables; uninitialized variables get the zero value (`""` for strings).
- The `+` operator concatenates strings.
- `:=` is the short variable declaration, inferring the type.
- `i++` is the only increment operator (there is no prefix form).
- Parentheses are not used in the `for` statement.
- Braces are required, and the opening brace must be on the same line as the `for`.
- Any part of the `for` statement (init, condition, post) can be omitted.

**Example:** Using a range-based for loop:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```

  - Each iteration of the `range` loop provides the index and value.
  - The blank identifier `_` is used to ignore the index when not needed.

Variable declaration styles:

 ```go
 s := ""             // short variable declaration
 var s string        // var declaration
 var s = ""          // var declaration with initialization
 var s string = ""   // var declaration with type and initialization
 ```

  - `:=` can only be used inside functions.
  - The second form uses the zero value for initialization.
  - The third and fourth forms are more explicit but less common in practice.

### 1.3 Funding Duplicate Lines

This section covers three variants of a program called `dup`, partially inspired by the Unix command `uniq`, which looks for adjacent duplicate lines.

**`Dup` Version 1**: prints each line that appears more than once in the input, preceeded by its count.

```go
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
```
[File: `dup1.go`](./ch01/duplicate_lines/dup1/dup1.go).

**Output:**  
<img src='images/20250426044014.png' width='450'/>

Things to note:
- A `map` holds a set of key-value pairs and provides constant-time access to the store.
- The key may be any type whose values  can be compared with `==`, strings being the most common.
- In this example, the key is a string and the value is an `int`.
- The build-in function creates a new empty map.
- The statement `counts[input.Text()]++` is a shorthand for the following:
```go
line := input.Text()
counts[line] = counts[line] + 1
```
- Each time `dup` reads a line of input, the line is used as a key in the map, and the value is incremented by 1.
- The `bufio` package provides Scanner, which is the easiest way to process input that naturally consists of lines.
- The `Scan` function returns `true` if there is another line to read, and `false` when the input is exhausted.
- The following table summarizes the format verbs used in `fmt.Printf`:

| Format     | Description                           |
| ---------- | ------------------------------------- |
| %d         | decimal integer                       |
| %x, %o, %b | integer in hexadecimal, octal, binary |
| %f         | floating-point number: 3.141593       |
| %g, %e     | 3.141592653589793  3.141593e+00       |
| %t         | boolean: true or false                |
| %c         | rune (Unicode code point)             |
| %s         | string                                |
| %q         | quoted string "abc" or rune 'c'       |
| %v         | any value in a natural format         |
| %T         | type of any value                     |
| %%         | literal percent sign (no operand)     |
- `printf` does not add a newline at the end of the output, so you need to add it manually with `\n`.


**`Dup` Version 2**: reads from files and counts duplicate lines.

```go
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
```
[File: `dup2.go`](./ch01/duplicate_lines/dup2/dup2.go).

**Output:**  
<img src='images/20250426052843.png' width='450'/>

Things to note:
- The function `os.Open` returns two values: a file pointer and an error value.
- If `err` is not `nil`, something went wrong while opening the file.
- The verb `%v` in `fmt.Printf` is used to print the value in a default format.
- The call to `countLines` precedes its declaraion; functions and other package-level declarations can be in any order.
- A `map` is a reference to a data structure created by the `make` function.
- When a `map` is passed to a function, the function receives a copy of the reference to the map, not a copy of the map itself.
- Any changes made to the map in the function are reflected in the original map.


The prior two versions of `dup` are not very efficient because input is read and broken into lines as needed. An alternative approach is to read the entire input into memory and then process it. This is done in the third version of `dup`.

**`Dup` Version 3**: reads from files and counts duplicate lines, but reads the entire input into memory first.

```go
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
```
[File: `dup3.go`](./ch01/duplicate_lines/dup3/dup3.go).

**Output:**  
<img src='images/20250426053007.png' width='450'/>

Things to note:
- `ReadFile` contains a byte slice that must be converted to a string before it can be split into lines.

**Exercise 1.4**: Modify the `dup2` program to print the names of all files in which each duplicated line occurs.  
[File: `dup4.go`](./ch01/duplicate_lines/ex1_4-dup4/dup4.go).

### 1.4 Animated GIFs

This next program demonstrates the use of Go's standard image packages to create *Lissajous figures*, which are parametric curves produced by harmonic oscillation in two dimensions.

New constructs introduced in this program include: const declarations, struct types, composite literals.

```go
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

// This block defines the color palette and constants for color indices
var palette = []color.Color{color.White, color.Black}   // Composite literal for color.Color slice

const (
    whiteIndex = 0
    blackIndex = 1
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of frames in the animation
        delay   = 8     // delay between frames in 10ms units
    )

    freq  := rand.Float64() * 3.0           // relative frequency of y oscillator
    anim  := gif.GIF{LoopCount: nframes}    // composite literal for gif.GIF struct
    phase := 0.0                            // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles * 2 * math.Pi;  t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)  // Note: using gif.EncodeAll instead of gif.Encode
}
```
[File: `lissajous.go`](./ch01/animated_gifs/lissajous_1/lissajous.go).

**Output:**  
![File: `lissajous.gif`](./ch01/animated_gifs/lissajous_1/out.gif).

Things to note:
- After importing a package whose path has multiple components, like `image/color`, you can refer to the package by its last component, like `color`, e.g. `color.White`.
- A `const` declaration declareas values that are fixed at compile time. `const` declarations may be used at the package level or inside functions.
- The expressions `[]color.Color{...}` and `gif.GIF{...}` are *composite literals*, a compact notation for instantiating any of Go's composite types, including structs, arrays, and slices.
- The type `gif.GIF` is a struct type, which is a collection of fields, each with a name and type. 

**Exercise 1.5**: Change the Lissajous program’s color palette to green on black, for added authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.

[File: `lissajous.go`](./ch01/animated_gifs/lissajous_ex1_5/lissajous.go).

**Output:**  
![File: `lissajous.gif`](./ch01/animated_gifs/lissajous_ex1_5/out.gif)

**Exercise 1.6**: Modify the Lissajous program to produce images in multiple colors by adding more values to palette and then displaying them by changing the third argument of SetColorIndex in some interesting way.

[File: `lissajous.go`](./ch01/animated_gifs/lissajous_ext1_6/lissajous.go).

**Output:**  
![File: `lissajous.gif`](./ch01/animated_gifs/lissajous_ext1_6/out.gif)

### 1.5 Fetching a URL

```go
package main

import (
    "fmt"
    "io/ioutil"
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
        // Read the response and store it in b
        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}
```
[File: `fetch.go`](./ch01/urls/fetch_1/fetch.go).

**Output:**  
<img src='images/20250505042912.png' width='550'/>

**Exercise 1.7**: The function call `io.copy(dst, src)` reads from src and writes to dst. use it instead of `ioutil.readall` to copy the response body to os.stdout without requiring a buffer large enough to hold the entire stream. be sure to check the error result of io.copy.

```go
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
```
[fetch.go](./ch01/urls/ex-1_7/fetch.go)

**Output:**  
<img src='images/20250505042912.png' width='550'/>

**Exercise 1.8**: Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.

```go
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
```
[fetch.go](./ch01/urls/ex-1_8/fetch.go)

**Exercise 1.9**: Modify fetch to also print the HTTP status code, found in resp.Status.

```go
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
```
[fetch.go](./ch01/urls/ex-1_9/fetch.go)

### 1.6 Fetching URLs Concurrently

Go's concurrency model is one of its more interesting features, allowing you to run multiple tasks simultaneously using goroutines and channels.

The next program, 'fetchall', fetches many URLs concurrently, so that the process will take no longer than the longest fetch rather than the sum of all fetch times. This version discards the responses but reports the size and elapsed time for each one:

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Runs concurrent fetches for each URL provided as a command-line argument and prints the results.
func main() {
	start := time.Now()
	ch    := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)           // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)           // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Fetches the content of a URL and sends the result or error to the provided channel.
func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)       // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()               // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
```
[fetchall.go](./ch01/concurrency/fetchall_1/fetchall.go)

**Output:**

```cmd
go build .\fetchall.go
.\fetchall.exe https://golang.org https://gopl.io https://godoc.org
0.19s   62937 https://golang.org
0.23s   33482 https://godoc.org
0.26s    4154 https://gopl.io
0.26s elapsed
```
Things to note:
- A *goroutine* is a concurrent function execution.
- A *channel* is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.
- The function `main` runs in a goroutine and the `go` statement creates additional goroutines.

**Exercise 1.10:** Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.

```go
	output := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Print(output)

	// Write the output to a file
	ioutil.WriteFile("fetch_output.txt", []byte(output), 0644)
```
[fetchall.go](./ch01/concurrency/ex1-10/fetchall.go)

**Output:** Note how the reported time hasn't changed, and the overall time is the same for both runs, indicating asynchronous fetching.
```cmd
.\fetchall.exe https://download.thinkbroadband.com/512MB.zip https://download.thinkbroadband.com/512MB.zip
0.62s     391 https://download.thinkbroadband.com/512MB.zip
0.62s     391 https://download.thinkbroadband.com/512MB.zip
0.62s elapsed
```

**Exercise 1.11:** Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in such cases.)

Using `fetchall` with longer argument lists results in a slight increase in the total elapsed time.

<img src="images/1749115257781.png" alt="fetchall performance graph" width="600">

### 1.7 A Web Server

Go's libraries make it easy to create a simple web server that respond to queries like those made by `fetch`. The following code is for a minimal server that returns the path component of the URL used to access the server.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) 	//each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))   // listens on port 8000
}

// handler echoes the Path component of the requested URL
func handler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)       // r is a struct representing the HTTP request with fields like URL, Method, Header, etc.
}
```
**Note:** Windows Defender may block the server from running, and you may need to add an exception this to work. See [Managing Windows Defender Exclusions](../../Notes/general.md#managing-windows-defender-exclusions).

**Output:**

<img src="images/1749200946580.png" alt="Server Output" width="500">

The following example adds a feature for returning the status:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far
func counter( w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
```


