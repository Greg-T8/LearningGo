<!-- omit in toc -->
# 2. Program Structure
- [2.1 Names](#21-names)
  - [2.2 Declarations](#22-declarations)
  - [2.3 Variables](#23-variables)
    - [2.3.1 Short Variable Declarations](#231-short-variable-declarations)
    - [2.3.2 Pointers](#232-pointers)
    - [2.3.3 The `new` Function](#233-the-new-function)
    - [2.3.4 Lifetime of Variables](#234-lifetime-of-variables)
  - [2.4 Assignments](#24-assignments)
    - [2.4.1 Tuple Assignment](#241-tuple-assignment)
    - [2.4.2 Assignability](#242-assignability)
    - [2.5 Type Declarations](#25-type-declarations)
    - [2.6 Packages and Files](#26-packages-and-files)
      - [2.6.1 Imports](#261-imports)


## 2.1 Names

The names of Go functions, variables, constants, types, statement labels, and packages follow a simple rule: a name begins with a letter or an underscore and may have any number of letters, digits, or underscores after that. 

Case matters: `heapSort` and `HeapSort` are different names.

As of this book's writing in 2015, Go has 25 reserved keywords: 

`break`, `case`, `chan`, `const`, `continue`, `default`, `defer`, `else`, `fallthrough`, `for`, `func`, `go`, `goto`, `if`, `import`, `interface`, `map`, `package`, `range`, `return`, `select`, `struct`, `switch`, `type`, and `var`.

The [Go Language Specification](https://go.dev/ref/spec?utm_source=chatgpt.com#Keywords) continues to list exactly these keywords.

In addition, there are about three dozen *predeclared* names like `int` and `true` for built-in constants, types, and functions:

| Category  | Predeclared Names                                                                                                                                                           |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Constants | `true` `false` `iota` `nil`                                                                                                                                                 |
| Types     | `int` `int8` `int16` `int32` `int64` `uint` `uint8` `uint16` `uint32` `uint64` `uintptr` `float32` `float64` `complex128` `complex64` `bool` `byte` `rune` `string` `error` |
| Functions | `make` `len` `cap` `new` `append` `copy` `close` `delete` `complex` `real` `imag` `panic` `recover`                                                                         |

These names are not reserved, so you may use them in declarations. There are a handful of scenarios where redeclaring one of them makes sense.

If an entity is declared within a function, it is *local* to that function and not visible outside of it. If it is declared at the package level, it is *global* to the package and may be used by any function in that package.

The case of the first letter of a name determines its visibility across package boundaries. If the name begins with an upper-case letter, it is *exported* and may be used by other packages, e.g. `Printf` in the `fmt` package. 

Package names are always lower-case.

There is no limit on name length, but convention and style in Go programs lean toward short names, especially for local variables and small scopes. Generally, the larger the scope of a name, the longer and more meaningful it should be.

Sylistically, Go programmers use "camel case" when forming names, i.e. interior capital letters are preferred over interior underscores. The letters of acronyms like ASCII and HTML are always rendered in the same case, e.g. `htmlParser`, `HTMLEscape`, or `escapeHTML`, but not `escapeHtml`.

### 2.2 Declarations

A *declaration* names a program entity and specifies some or all of its properties.

There are four major kinds of declarations: `var`, `const`, `type`, and `func`.

A Go program is stored in one or more files with the `.go` extension. Each file begins with a package declaration that says what package the file is part of.

The package declaration is followed by any import declarations and then a sequence of *package-level* declarations of types, varaibles, constants, and functions.

```go
// Boiling prints the boiling point of water.
package main					// package-level declaration

import "fmt"

const boilingF = 212.0			// package-level declaration

func main() {
    f := boilingF				// local declaration
    c := (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g degrees Fahrenheit or %g degrees Celsius\n", f, c)
}
```

```go
// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%g degrees Fahrenheit = %g degrees Celsius\n", freezingF, fToC(freezingF))
    fmt.Printf("%g degrees Fahrenheit = %g degrees Celsius\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
    return (f - 32) * 5 / 9
}
```

### 2.3 Variables

A `var` declaration creates a variable of a particular type, attaches a name to it, and sets its initial value.

```go
var name type = expression
```

Either the type or the `= expression` part may be omitted, but not both. 

If the type is omitted, it is inferred from the expression. 

If the expression is omitted, the initial value is the *zero value* for the type, which is `0` for numeric types, `false` for booleans, and `""` for strings. The zero value of an aggregate type like an array or struct is the zero value of each of its elements.

The zero-value mechanism ensures a variable always holds a well-defined value of its type; *in Go there is no such thing as an uninitialized variable.*

```go
var s string        // declares a variable s of type string, initialized to the zero value ""  
fmt.Println(s)      // prints an empty string rather than an error
```

You can declare and optionally initialize multiple variables in a single `var` declaration:

```go
var i, j, k int                     // declares three variables of type int, all initialized to 0
var b, f, s = true, 2.3, "four"         // declares and initializes three variables of different types, bool, float64, and string
```

Package-level variables are initialized before the `main` function is called. Local variables are initialized as their declarations are encountered during function execution.

A set of variables can be initialized by calling a function that returns multiple values:

```go
var f, err = os.Open(name)      // declares two variables f and err, initialized by the return values of os.Open(name)
```

#### 2.3.1 Short Variable Declarations

Within a function, an alternate form called a *short variable declaration* may be used to declare and inintialize local variables:

```go
name := expression
```

The type of `name` is inferred from the type of `expression`. 

Behind the scenes, this is like writing:

```go
var name type = expression
```
Examples:

```go
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64() * 3.0
t    := 0.0
```

Short-variable declarations are used to declare and initialize the majority of local variables.

A `var` declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression, or for when the variable will be assigned a value later and its initial value is not important.

```go
i := 100                        // short variable declaration, i is of type int
var boiling float64 = 100       // var declaration with initialization, boiling is of type float64

var names []string              // declares a variable names of type slice of string, initialized to the zero value nil
var err error
var p Point
```

In a short variable declaration, multiple variables may be declared and initialized in a single statement:

```go
i, j := 0, 1
```

However, declarations with mulitiple initializer expressions should only be used when they help readability, such as for short and natural groupings like the initialization part of a `for` loop.

Keep in mind that `:=` is a declaration, whereas `=` is an assignment&mdash; `=` does not declare a new variable, it *requires* the variable already declared.

Short-variable declarations may be used for calls to functions like `os.Open` that return multiple values:

```go
f, err := os.Open(name)
if err != nil {
    return err
}
// ...use f...
f.Close()
```

**Important:** A short variable declaration does not necessarily *declare* all the variables on the left-hand side. If some of them were already declared in the same lexical block, then the short variable declaration acts as an assignment to those variables.

```go
in, err := os.Open(infile)         
// ...
out, err := os.Create(outfile)     // out is a new variable, but err is the same as before
```

A short variable declaration must declare at least one new variable, so the following code won't compile:

```go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile)        // error: no new variables on left side of := operator
```

To fix this, you can use an ordinary assignment for the second statement:

```go
f, err := os.Create(outfile)        
// ...
f, err = os.Create(outfile)        // ordinary assignment, reusing f and err
```

#### 2.3.2 Pointers

A *variable* is a piece of storage containing a value.

Variables created by declarations are identified bya name, such as `x`, but many variables are identified only by expressions like `x[i]` or `x.f`.

All of these expressions read the value of a variable, except when they appear on the left-hand side of an assignment, in which case a new value is assigned to the variable.

A *pointer* value is the *address* of a variable. A pointer is thus the location at which a value is stored. Not every value has an address, but every variable does.

With a pointer, you can read or update the value of a variable *indirectly* without using or even knowing the name of the variable.

If you declare a variable like `var x int`, then using `&x` gives you the memory address where `x` is stored. This address is represented by a value of type `*int`, which means "a pointer to an integer."

```go
x := 1
println(&x)			// Using &x to get the address of x, i.e. printing the address of x
```
If you assign this address to a variable named `p`, you can say that "p points to x" — in other words, `p` holds the memory address where `x` is stored.

```go
x := 1
p := &x 	        // p is a pointer to x, i.e. p holds the address of x
```
To access the value that `p` points to, you write `*p`. This gives you the actual integer stored at that memory address. 

```go
println(*p)		    // Using *p to dereference p, i.e. printing the value at the address stored in p
```

Since `*p` refers to a variable, you can also use it on the left side of an assignment to change its value — that is, update the variable that `p` points to.

```go
*p = 2			    // Using *p to change the value at the address stored in p
```

```go
x := 1
p := &x				// p, of type *int, points to x
fmt.Println(*p)		// "1"
*p = 2				// equivalent to x = 2
fmt.Println(x)		// "2"
```

When you have a variable made up of multiple parts — like the fields in a `struct` or items in an array — each part is its own variable and also has its own memory address.

Variables are sometimes called “addressable values” because they occupy a specific location in memory. Only expressions that refer to variables can be used with the `&` operator to get their memory address.

By default, a pointer that hasn't been assigned a value is `nil`, meaning it doesn't point to anything. You can check if a pointer `p` is valid by testing `p != nil`. This will be true if `p` actually points to a variable.

You can also compare pointers: two pointers are considered equal if they both point to the same memory location — or if they’re both `nil`.

```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil)	// "true false false"
```

It is perfectly safe for a function to return the address of a local variable. In the following code, the local variable `v` created by this call to `f()`will remain in existence even after the call has returned, and the pointer `p` will still refer to it.

```go
var p = f()         // p is a pointer to the local variable v created by f()

func f() *int {     // f returns a pointer to a local variable
    v := 1
    return &v       // returns the address of v
}
```

Each call of `f()` returns a distinct value:

```go
fmt.Println(f() == f())     // "false"
```

A pointer holds a variable’s address. So, when you pass a pointer to a function, the function can modify the original variable indirectly.

```go
func incr(p *int) {
    *p++				// increments what p points to; does not change p, which still points to the same variable
    return *p           // *p is an alias for v
}

v := 1
incr(&v)				// side effect:  v is now 2
fmt.Println(incr(&v))	// "3" (and v is 3)
```

Each time we take the address of a variable or copy a pointer, we create new aliases or ways to identify the same variable. For example, `*p` is an alias for `v`. 

Pointer aliasing is useful because it allows us to access a variable without using its name. But, this is a double-edged sword: to find all the statements that access a variable, we have to know all its aliases.

It's not just pointers that create aliases; aliasing also occurs when we copy values of other reference types, like slices, maps, channels, structs, arrays, and interfaces.

Pointers are key to the `flag` package, which uses a program's command-line arguments to set the values of certain variables distributed throughout the program. 

```go
// Echo4 prints its command-line arguments.
package main

import (
    "flag"
    "fmt"
    "strings"
)

var (
    // Optional command-line flags to control the output.

    // flag.Bool() arguments: <name>, <default value>, <help message>; returns a pointer to a bool
    n = flag.Bool("n", false, "omit trailing newline")

    // flag.String() arguments: <name>, <default value>, <help message>; returns a pointer to a string
    sep = flag.String("s", " ", "separator")
)

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep)) // Non-flag arguments are accessed via flag.Args();
    if !*n {                                   // n is a pointer to a bool, so we dereference it with *
        fmt.Println()
    }
}
```

```pwsh
go build .\echo4.go

.\echo4.exe a bc def                
a bc def

.\echo4.exe -s / abc def            # Using the -s flag to set a custom separator
abc/def

.\echo4.exe -n a bc def             # Using the -n flag to omit the trailing newline
a bc def                            # Note: no newline at the end

.\echo4.exe -help
Usage of C:\Users\gregt\LocalCode\LearningGo\Books\the_go_programming_language\ch02\echo4.exe:
  -n    omit trailing newline
  -s string
        separator (default " ")
```

#### 2.3.3 The `new` Function

Another way to create a variable is to use the built-in `new` function. The expression `new(T)` creates an *unnamed variable* of type `T`, initializes it to the zero value of `T`, and returns its address, which is a value of type `*T`.

```go
p := new(int)		// p, of type *int, points to an unnamed int variable
fmt.Println(*p)		// "0" is printed, dereferencing p to get the value it points to
*p = 2				// sets the unnamed int to 2
fmt.Println(*p)		// "2" is printed, dereferencing p again
```

A variable created with `new` works just like a regular local variable whose address is taken. The difference is that `new` doesn't require you to create and name a dummy variable. 

You can also use `new(T)` directly in an expression. So, `new` is just a convenience feature, not a core language concept. For example, the two `newInt` functions below behave the same way, even if one uses `new` and the other doesn’t.

```go
func newInt() *int {
    var dummy int           // Using a named int variable
    return &dummy
}

func newInt() *int {
    return new(int)         // Convenience: creates an unnamed int variable and returns its address
}
```
Each call to `new()` returns a distinct variable with a unique address:

```go
p := new(int)
q := new(int)
fmt.Println(p == q)		    // false, p and q point to different memory addresses
```

There is one exception to this rule: two variables whose type carries no information and is therefore of size zero, such as `struct{}` or `[0]int`, may, depending on the implementation, have the same address.

The `new` function is rarely used because the most common unnamed variables are of struct types, for which the `struct` literal syntax is more flexible.

Since `new` is a predeclared function, not a keyword, it is possible to redefine the meaning for something else within a function:

```go
func delta(old, new int) int { return new - old }           // Redefining new, as it is a predeclared function, not a keyword
```

This means the built-in `new` function is not available in the scope of the `delta` function.

#### 2.3.4 Lifetime of Variables

A variable’s lifetime is the period during program execution when it exists.

Package-level variables last for the entire program run.

Local variables have dynamic lifetimes: a new instance is created each time the declaration is executed and remains until it becomes unreachable, after which its memory can be reused.

Function parameters and return values are also local variables, created each time the function is called.
For example, in the Lissajous program from Section 1.4:

```go
for t := 0.0; t < cycles*2*math.Pi; t += res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
}
```

Here, `t` is created at the start of each loop execution.
`x` and `y` are created on each iteration.

The garbage collector reclaims a variable’s memory when no active reference path exists from any package-level variable or any local variable in a currently running function.
If no such path exists, the variable is unreachable and can be discarded.

Because lifetime depends on reachability, a local variable may survive beyond a single loop iteration or even after its enclosing function returns.

A compiler may allocate local variables on either the heap or the stack, and this choice is not determined by whether they were declared with `var` or `new`.

```go
var global *int

func f() {
    var x int
    x = 1
    global = &x             // x escapes from f, so it must be heap-allocated
}

func g() {                  
    y := new(int)           // y does not escape from g, so it can be stack-allocated
    *y = 1                  
}
```

In `f`, `x` must be heap-allocated because it remains reachable through `global` after `f` returns—`x` escapes from `f`.
In `g`, `*y` becomes unreachable once the function returns, so it can be recycled. Since `*y` does not escape, the compiler can place it on the stack, even though it was created with `new`.

Escape analysis is not required for writing correct programs, but it matters for performance—variables that escape require extra memory allocation.

Garbage collection makes writing correct code easier, but you still need to think about memory.

Holding references to short-lived objects inside long-lived objects (especially globals) can prevent the garbage collector from reclaiming them.

### 2.4 Assignments

An assignment statement updates a variable’s value. In its simplest form, the variable is on the left of `=` and an expression is on the right.

```go
x = 1                               // named variable
*p = true                           // indirect variable
person.name = "bob"                 // struct field
count[x] = count[x] * scale         // array, slice, or map element
```

Each arithmetic and bitwise binary operator has an assignment form, so the last example can be written as:

```go
count[x] *= scale
```

This avoids repeating (and re-evaluating) the variable expression.

Numeric variables can also be updated with increment (`++`) and decrement (`--`):

```go
v := 1
v++         // same as v = v + 1; v becomes 2
v--         // same as v = v - 1; v becomes 1
```

#### 2.4.1 Tuple Assignment

Tuple assignment lets you assign values to multiple variables in a single statement. All right-hand side expressions are evaluated before any variables are updated, which is useful when variables appear on both sides, such as in swaps:

```go
x, y = y, x                     // swap the values of x and y
a[i], a[j] = a[j], a[i]         // swap the elements of a at indices i and j
```

It’s also handy in algorithms like computing the greatest common divisor:

```go
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y           // swap x and y, then set y to the remainder of x divided by y
    }
    return x
}
```

Or for calculating the n-th Fibonacci number iteratively:

```go
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y           // swap x and y, then set y to the sum of x and y
    }
    return x
}
```

Tuple assignment can make trivial assignments compact:

```go
i, j, k = 2, 3, 5
```

However, if the expressions are complex, use separate statements for readability.

Some expressions produce multiple results, such as function calls. In that case, the left-hand side must have the same number of variables:

```go
f, err = os.Open("foo.txt")         // returns two values
```

Often, the extra value signals an error or a boolean “ok” flag. The same pattern applies to:

```go
v, ok = m[key]          // map lookup
v, ok = x.(T)           // type assertion
v, ok = <-ch            // channel receive
```

Unwanted values can be discarded with the blank identifier:

```go
_, err = io.Copy(dst, src)      // ignore byte count
_, ok = x.(T)                   // check type, discard result
```

#### 2.4.2 Assignability

Assignment statements are explicit assignments, but assignments also happen implicitly in several cases:

* A function call assigns argument values to the corresponding parameters.
* A return statement assigns the return values to the corresponding result variables.
* A composite literal assigns values to its elements. For example:

```go
medals := []string{"gold", "silver", "bronze"}
```

is equivalent to:

```go
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```

Elements of maps and channels, though not standard variables, also receive values through implicit assignment.

An assignment—explicit or implicit—is valid if the variable’s type and the value’s type match, or if the value is assignable to the variable’s type. For types covered so far, this means the types must match exactly, with `nil` allowed for interface or reference types.

Constants have more flexible assignability rules, often avoiding explicit conversions.

Comparability (`==` and `!=`) is related to assignability: in a comparison, one operand must be assignable to the other’s type.

#### 2.5 Type Declarations

A variable’s type defines what values it can hold, how those values are stored, what operations are valid, and which methods are available.

Different variables may share the same underlying representation but mean very different things. For example, an `int` might store a loop index, a timestamp, a file descriptor, or a month. A `float64` could represent speed or temperature, and a `string` could be a password or a color name.

A type declaration creates a new named type based on an existing type:

```go
type name underlying-type
```

Named types have the same underlying type as the original but are treated as distinct. This prevents unintended mixing of incompatible concepts. If the name is exported (starts with an uppercase letter) and declared at package level, it is available outside the package.

**Example** — defining Celsius and Fahrenheit as distinct types:

```go
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

Although both `Celsius` and `Fahrenheit` use `float64` underneath, they cannot be compared or combined directly. Conversions like `Celsius(t)` or `Fahrenheit(t)` make the meaning explicit and do not change the stored value.

The functions `CToF` and `FToC` actually change the value to represent a different scale.

For any type `T`, `T(x)` converts `x` to `T`. Conversions are allowed if both types share the same underlying type, or if both are unnamed pointer types pointing to the same underlying type. These conversions change the type but not the representation.

Conversions are also allowed between numeric types, and between strings and certain slices. These may change the representation — for example, converting a float to an integer drops the fractional part, and converting a string to a `[]byte` makes a copy of its data.

All type conversions succeed at runtime.

The underlying type of a named type defines its structure, representation, and the intrinsic operations it supports — exactly as if the underlying type were used directly.

For example, arithmetic works for `Celsius` and `Fahrenheit` just like for `float64`:

```go
fmt.Printf("%g\n", BoilingC - FreezingC)          // "100" °C, %g prints a floating-point number, automatically choosing the best format, i.e. decimal or scientific notation
boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF - CToF(FreezingC))    // "180" °F
fmt.Printf("%g\n", boilingF - FreezingC)          // compile error: type mismatch
```

Comparison operators also work between values of the same named type, or between a named type and an unnamed type with the same underlying type. But you can’t directly compare two different named types:

```go
var c Celsius
var f Fahrenheit
fmt.Println(c == 0)                 // "true"
fmt.Println(f >= 0)                 // "true"
fmt.Println(c == f)                 // compile error: type mismatch
fmt.Println(c == Celsius(f))        // "true"
```

Here, `Celsius(f)` doesn’t change `f`’s value, only its type. Both `c` and `f` are zero, so the comparison is true.

Named types are more useful when the underlying type is complex, as they avoid repeating long type expressions. They also allow defining methods for added behavior.

For example, a `String` method for `Celsius`:

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

The `String` method controls how values print when used with `fmt`:

```go
c := FToC(212.0)
fmt.Println(c.String())         // "100°C"
fmt.Printf("%v\n", c)           // "100°C", %v means "default format": it prints the value based on its type
fmt.Printf("%s\n", c)           // "100°C", %s means "string format": interprets the value as a string
fmt.Println(c)                  // "100°C"
fmt.Printf("%g\n", c)           // "100" — does not call String
fmt.Println(float64(c))         // "100" — does not call String
```

#### 2.6 Packages and Files

In Go, packages serve the same role as libraries or modules in other languages. They provide modularity, encapsulation, separate compilation, and reuse. A package’s source code is stored in one or more `.go` files, typically within a directory whose name matches its import path. For example, the files for the `gopl.io/ch1/helloworld` package are found in `$GOPATH/src/gopl.io/ch1/helloworld`.

Each package is its own namespace. For instance, `image.Decode` and `utf16.Decode` are different functions, even though they share the same name. To use a function from another package, you must qualify it with the package name.

Packages also control visibility. Identifiers that begin with an uppercase letter are exported and available outside the package; those starting with a lowercase letter remain internal.

Suppose our temperature conversion program becomes popular and we want to share it as a package. We can create `gopl.io/ch2/tempconv`, a variation of our earlier example. Although a small package like this would usually need just one file, here we’ll split it into two files to show how declarations across multiple files in a package work. In `tempconv.go`, we define the types, constants, and methods:

[tempconv/tempconv.go]()
```go
// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }           // Sprintf returns a formatted string without printing
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
```
and the conversion functions in conv.go:

[tempconv/conv.go]()
```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

Each file begins with a package declaration that defines the package name. When imported, its members are accessed with the package qualifier, such as `tempconv.CToF`.

All package-level names, like types and constants, are shared across files in the same package as if they were in a single source file. For example, `tempconv.go` imports `fmt`, but `conv.go` does not, since it doesn’t need it.

Because the constants are exported (they start with uppercase letters), they can be accessed directly:

```go
fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
```

To convert Celsius to Fahrenheit in another package that imports `gopl.io/ch2/tempconv`, you can write:

```go
fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
```

A package-level doc comment should be placed immediately before the package declaration. It typically starts with a one-sentence summary of the package. Only one file in the package should contain this comment. Larger documentation is often put in a dedicated `doc.go` file.

**Exercise 2.1:** Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.

- [tempconv/tempconv.go](ch02/2.6/tempconv/tempconv.go)
- [tempconv/conv.go](ch02/2.6/tempconv/conv.go)

##### 2.6.1 Imports

Within a Go program, each package is identified by an import path, the string that appears in an `import` declaration such as `"gopl.io/ch2/tempconv"`. The Go specification doesn’t define the meaning of these strings; tools like the `go` command interpret them. With the `go` tool, an import path corresponds to a directory containing one or more `.go` files that form the package.

Besides its import path, every package has a package name, declared at the top of its files. By convention, this name is the last segment of the import path. For example, the package name for `gopl.io/ch2/tempconv` is `tempconv`.

To use it, we import the package in another program:

```go
// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
    "fmt"
    "os"
    "strconv"      

    "gopl.io/ch2/tempconv"                                  // Importing the tempconv package
)

func main() {
    for _, arg := range os.Args[1:] {                       // os.Args[1:] skips the program name
        t, err := strconv.ParseFloat(arg, 64)      
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)         // %v means "default format": it prints the value based on its type
            os.Exit(1)                                      // Fprintf formats a string and writes it to the specified output, here os.Stderr
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s\n",                    // %s means "string format": interprets the value as a string
            f, tempconv.FToC(f), c, tempconv.CToF(c))
    }
}
```

The import binds a short name to the package, which is used to access its members, like `tempconv.CToF`. By default, this short name matches the package name, but an alternative name can be specified to avoid conflicts, e.g.:
```go
import tc "gopl.io/ch2/tempconv"  // Importing tempconv as tc
```

The `cf` program converts a numeric command-line argument into both Celsius and Fahrenheit:

```bash
$ go build gopl.io/ch2/cf
$ ./cf 32
32°F = 0°C, 32°C = 89.6°F
$ ./cf 212
212°F = 100°C, 212°C = 413.6°F
$ ./cf -40
-40°F = -40°C, -40°C = -40°F
```

Go reports an error if a package is imported but never used. This prevents unnecessary dependencies but can be inconvenient while debugging. For example, removing the only use of `log` might leave an unused import, which then causes a compile error. In that case, the import must be removed or commented out.

A better solution is to use the `goimports` tool (`golang.org/x/tools/cmd/goimports`). It automatically adds and removes imports as needed and formats the code in Go’s standard style, just like `gofmt`. Most editors can be set up to run `goimports` on save.
