
## 10. The `go` Tool

This chapter focuses on the go tool, which handles downloading, querying, formatting, building, testing, and installing Go packages. It combines many functions into a single command set. As a package manager, it answers queries about installed packages, resolves dependencies, and fetches code from remote version-control systems. As a build system, it tracks file dependencies and runs compilers, assemblers, and linkers, though it is intentionally simpler than the Unix make utility. It also serves as a test driver, which will be discussed in Chapter 11.

The command-line interface follows a “Swiss army knife” style, with many subcommands. Some you have already seen include get, run, build, and fmt. Running `go help` shows the full documentation, but the most common commands are:

```bash
$ go
    build       compile packages and dependencies
    clean       remove object files
    doc         show documentation for package or symbol
    env         print Go environment information
    fmt         run gofmt on package sources
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    version     print Go version
    vet         run go tool vet on packages
```

For details on a command, run `go help [command]`.

The go tool minimizes the need for configuration by relying on conventions. For example, since each directory holds a single package, the tool can determine a package from a source file’s directory. The import path matches the directory hierarchy in the workspace, so the tool can locate both object files and the corresponding source code repository.
