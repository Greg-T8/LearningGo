# My Notes from "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan

<img src='images/20250419135706.png' width='300'/>

**Book Resources:**
- [Official Book Website](https://www.gopl.io/)
- [Source Code Examples](https://github.com/adonovan/gopl.io)


<!-- omit in toc -->
## Go Commands 

```pwsh
go
    build       # compile packages and dependencies
    clean       # remove object files
    doc         # show documentation for package or symbol
    env         # print Go environment information
    fmt         # run gofmt on package sources
    get         # download and install packages and dependencies
    install     # compile and install packages and dependencies
    list        # list packages
    run         # compile and run Go program
    test        # test packages
    version     # print Go version
    vet         # run go tool vet on packages

go run hello.go             # Run the program hello.go
go fmt hello.go             # Format the code in hello.go
go build hello.go           # Build the program hello.go into an executable file
go clean -cache             # Clean the Go build cache
go build -x -v hello.go     # Build the program with verbose output
go doc http.Get             # Show documentation for the http.Get function
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

## Go Programming Language Overview

Go is an open-source programming language designed to make building software simple, reliable, and efficient. It was created in September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google, and officially announced in November 2009. The language and its tools were built with three main goals: expressiveness, efficiency in both compilation and execution, and support for writing reliable, robust programs.

Go resembles C on the surface and, like C, is aimed at professional programmers who want powerful results with minimal complexity. However, it is more than a modernized C. It adapts successful ideas from many languages while avoiding features that create unnecessary complexity or unreliable code. Its concurrency model is both new and efficient, and its approach to data abstraction and object-oriented programming is unusually flexible. Automatic memory management through garbage collection is built in.

Go is especially suited for infrastructure such as network servers and developer tools, but it is a true general-purpose language. It is used in areas like graphics, mobile apps, and machine learning. Many developers adopt Go as a safer and faster alternative to untyped scripting languages, since Go programs typically perform better and avoid type-related runtime crashes.

As an open-source project, Go makes its compiler, libraries, and tools freely available, with contributions from a global developer community. It runs on Linux, FreeBSD, OpenBSD, macOS, Plan 9, and Windows. Programs written in one environment usually run on the others without modification.

### The Origins of Go

Like biological species, programming languages evolve. Successful ones pass on useful traits, hybrids gain unexpected strengths, and sometimes entirely new features appear. By tracing these influences, we can see why Go looks the way it does and what environments it was designed for.

Go is often called a “C-like language” or “C for the 21st century.” From C, it inherits expression syntax, control-flow statements, basic data types, call-by-value parameter passing, pointers, and a focus on generating efficient machine code that works well with operating system abstractions. But C is only part of its heritage. Niklaus Wirth’s languages also shaped Go. Modula-2 inspired its package system. Oberon removed the split between module interfaces and implementations. Oberon-2 influenced package and import syntax, as well as method declarations.
  
<img src='./images/20250419140256.png' width='400'/>

**Influencers of Go**:  
  - **C**: Syntax, control structures, basic data types, call-by-value parameter passing, pointers, and emphasis on compiling to machine code.
  - **Modula-2**: Inspired the package concept.
  - **Oberon**: Eliminated the distinction between module interface files and module implementation files.
  - **Oberon-2**: Influenced syntax for packages and declarations, particularly method declarations.
  - **CSP (Communicating Sequential Processes)**: Influenced goroutines and channels.
  - **Squeak**: Provided handling for mouse and keyboard events, with statically created channels.
  - **Newsqueak**: A purely functional language with garbage collection, aimed at managing keyboard, mouse, and window events.
  - **Alef**: Attempted to make Newsqueak a viable system programming language, but its lack of garbage collection made concurrency too painful.

Another major influence comes from research languages at Bell Labs, based on Tony Hoare’s 1978 paper on communicating sequential processes (CSP). CSP describes concurrency as independent processes that don’t share state but instead communicate and synchronize over channels. Though CSP was not a programming language itself, Rob Pike and others built languages around its ideas. Squeak was the first, handling mouse and keyboard events with statically created channels. Newsqueak followed, adding C-like syntax, Pascal-style types, garbage collection, and first-class channels that could be dynamically created and stored in variables. Plan 9’s Alef tried to extend these ideas for systems programming but dropped garbage collection, making concurrency too difficult.

Go also borrows from other places. The iota constant generator resembles APL. Nested functions with lexical scope come from Scheme and other modern languages. At the same time, Go introduces new ideas of its own. Slices provide dynamic arrays with efficient indexing and flexible sharing, similar to linked structures. The defer statement, another original feature, adds a new way to manage cleanup and resource handling.



### The Go Project

Programming languages reflect the philosophies of their creators, often reacting to the weaknesses of earlier languages. Go was created out of frustration with the growing complexity of large Google systems, a problem common across the industry. As Rob Pike observed, “complexity is multiplicative”: fixing one problem by adding complexity in one place inevitably spreads complexity elsewhere. With constant pressure to add features, ship quickly, and provide endless configuration options, simplicity is often overlooked—even though it is the foundation of good software. True simplicity requires more effort at the start of a project and discipline throughout its life to preserve what Fred Brooks called “conceptual integrity.” Good changes can be integrated without harming the design, but bad ones undermine it, and harmful changes trade lasting simplicity for shallow convenience. A system stays stable, secure, and coherent only through simple design.

The Go project reflects this belief. It consists of the language, its tools and libraries, and a culture of radical simplicity. Go benefits from hindsight: it has garbage collection, a package system, first-class functions, lexical scope, system call interfaces, and UTF-8–encoded immutable strings. Yet it has relatively few features and is unlikely to add more. It omits implicit numeric conversions, constructors and destructors, operator overloading, default parameters, inheritance, generics, exceptions, macros, annotations, and thread-local storage. Go is mature, stable, and guarantees backward compatibility—older programs continue to work with newer compilers and libraries.

Go’s type system is strong enough to catch many common mistakes from dynamic languages but is simpler than those in C++ or Haskell. This sometimes results in limited “untyped” code, but in practice Go provides most of the safety and performance of stronger type systems without their complexity. The language also reflects modern hardware concerns, especially locality. Its built-in types and library structures work naturally without hidden allocations, implicit constructors, or unnecessary indirection. Structs and arrays store elements directly, reducing memory use and pointer overhead.

Concurrency is central to Go’s design. Goroutines are lightweight threads with small, variable-size stacks. Starting one is cheap, and creating millions is feasible.

Go’s standard library, often described as having “batteries included,” supplies clean, conventional building blocks for I/O, text processing, graphics, cryptography, networking, distributed systems, and many standard file formats and protocols. Libraries and tools emphasize convention over configuration, reducing complexity and making Go projects easier to learn and maintain. The `go` tool relies only on filenames, identifiers, and a few special comments to infer a project’s libraries, executables, tests, benchmarks, examples, platform-specific variants, and documentation—the Go source itself serves as the build specification.
