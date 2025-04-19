# My Notes from "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan

<img src='images/20250419135706.png' width='350'/>

<details>
<summary>Book Resources</summary>

- [Official Book Website](https://www.gopl.io/)
- [Source Code Examples](https://github.com/adonovan/gopl.io)

</details>

## Go Commands 

```Go
go run hello.go         // Run the program hello.go
go build hello.go       // Build the program hello.go into an executable file

```


## Overview and History of Go

<details>
<summary>Overview and History of Go</summary>

- **Conception**: Go was conceived in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google and was publicly announced in 2009.

- **Influencers of Go**:
  
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

</details>

## Hello World

The following program can be compiled and ran with `go run hello.go`:

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

The command `go build hello.go` creates an executable file.

## Go Basics

- Go code is organized into packages, which are similar to libraries or modules in other languages.
- Each source file begins with a package declaration, which states the package the file belongs go.

```Go
package main
```

- You must import *exactly* the packages you need. The Go compiler will not compile a package if it is not used in the code.
- Go does not require semicolons; they are automatically inserted at the end of lines.

### Command-Line Arguments

**Slices**:
- The variable `os.Args` is a slice of strings.
- The first element of `os.Args`, `os.Args[0]`, is the name of the command.
- The other elements, `os.Args[1:len(os.Args)]`, are the command-line arguments passed to the program.

- The following program mimics echos the command-line arguments passed to it:

```Go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize empty string variables for the result and separator
	var s, sep string
	// Iterate over command-line arguments (excluding the program name)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// Print the concatenated arguments
	fmt.Println(s)
}
```

See [02-echo1.go](./ch01/02-echo1.go) for the complete code.

Things to note:
- The `var` declaration declares two variables, `s` and `sep`. 
- If the variable is not explicitly initialized, it is implicitly initialized to the zero value of its type. For strings, the zero value is an empty string `""`.
- The `+` operator concatenates strings.
- 