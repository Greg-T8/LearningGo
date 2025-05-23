# My Notes from "Learning Go" by Jon Bodner
<img src='images/20250413052634.png' width='250'/>

<details>
<summary>Book Resources</summary>

- [Book Code Examples](https://github.com/learning-go-book-2e)
- [Book Website](https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/)

</details>

## Go Language Resources
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Wiki: Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Go in VSCode](https://code.visualstudio.com/docs/languages/go)
- [Delve Debugger](https://github.com/go-delve/delve/)

## Go Commands

```powershell
go version                      # Check Go version
go mod init hello_world         # Initialize a new module
go fmt ./...                    # Format all Go files in the current directory
go vet ./...                    # Check for potential issues in the code
go build                        # Build the current package
go build -o hello.exe           # Build and name the output file to another name
go install github.com/go-delve/delve/cmd/dlv@latest # Install Delve debugger
```

## General Notes

### The Semicolon Insertion Rule
- Go developers should never put semicolons at the end of lines, as Go inserts them automatically.
```Go
// Correct
func main()
{
    fmt.Println("Hello, World!")
}
// Incorrect - semicolon at the end of the line results in two semicolons; not valid Go
func main(); {
    fmt.Println("Hello, World!");
}
```

Use `go fmt` to format your code. It automatically inserts semicolons where needed.

### Predeclared Types

- Predeclared types are types that are built into the Go language. 
- Here's the list of the predeclared types:
 




  - The Zero Value: Go assigns a default *zero* value to variables that are declared but not initialized. See [The zero value](https://go.dev/ref/spec#The_zero_value)
