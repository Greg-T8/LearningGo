# My Notes from "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan

<img src='images/20250419135706.png' width='350'/>

<details>
<summary>Book Resources</summary>

- [Official Book Website](https://www.gopl.io/)
- [Source Code Examples](https://github.com/adonovan/gopl.io)

</details>

<!-- omit in toc -->
## Go Commands 

```cmd
go
    build       // compile packages and dependencies
    clean       // remove object files
    doc         // show documentation for package or symbol
    env         // print Go environment information
    fmt         // run gofmt on package sources
    get         // download and install packages and dependencies
    install     // compile and install packages and dependencies
    list        // list packages
    run         // compile and run Go program
    test        // test packages
    version     // print Go version
    vet         // run go tool vet on packages

go run hello.go                         // Run the program hello.go
go fmt hello.go                         // Format the code in hello.go
go build hello.go                       // Build the program hello.go into an executable file
go clean -cache                         // Clean the Go build cache
go build -x -v hello.go                 // Build the program with verbose output
go doc http.Get                         // Show documentation for the http.Get function
```
<!-- omit in toc -->
## My Notes

- [Chapter 1: Tutorial](ch01/notes.md)
  - [1.2 - Command-Line Arguments](ch01/notes.md#12-command-line-arguments)
  - [1.3 - Finding Duplicate Lines](ch01/notes.md#13-funding-duplicate-lines)
  - [1.4 - Animated Gifs](ch01/notes.md#14-animated-gifs)
  - [1.5 - Fetching a URL](ch01/notes.md#15-fetching-a-url)
  - [1.6 Fetching URLs Concurrently](ch01/notes.md#16-fetching-urls-concurrently)
  - [1.7 A Web Server](ch01/notes.md#17-a-web-server)
  - [1.8 Loose Ends](ch01/notes.md#18-loose-ends)
- [Chapter 2: Program Structure](ch02/notes.md)
  - [2.1 Names](ch02/notes.md#21-names)
  - [2.2 Declarations](ch02/notes.md#22-declarations)
  - [2.3 Variables](ch02/notes.md#23-variables)
    - [2.3.1 Short Variable Declarations](ch02/notes.md#231-short-variable-declarations)
    - [2.3.2 Pointers](ch02/notes.md#232-pointers)
    - [2.3.3 The `new` Function](ch02/notes.md#233-the-new-function)
    - [2.3.4 Lifetime of Variables](ch02/notes.md#234-lifetime-of-variables)
    - [2.4.1 Tuple Assignment](ch02/notes.md#241-tuple-assignment)
    - [2.4.2 Assignability](ch02/notes.md#242-assignability)
  - [2.5 Type Declarations](ch02/notes.md#25-type-declarations)
  - [2.6 Packages and Files](ch02/notes.md#26-packages-and-files)
    - [2.6.1 Imports](ch02/notes.md#261-imports)
- [Chapter 10: The `go` Tool](ch10/notes.md)

## Overview and History of Go

**Conception**: Go was conceived in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google and was publicly announced in 2009.

**Influencers of Go**:
  
  <img src='./images/20250419140256.png' width='400'/>

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
