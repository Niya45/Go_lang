# Packages and Modules

```go
package main
func main() {}
```

- **Package** : a named collection of Go files. Multiple packages in the same dir aren't allowed 
    - `package main` : defines an executable program (has a main funct)
    - `import "fmt"` : import built in packages 
- **Module** : a collection of two or more packages
    - `go mod init [directory]` : makes the dir a module
    - `go.mod` : contains name of module version of go and external module

---

- `go run file.go` : creates a temporary binary file and executes the file, deleting the binary after execution
- `go build file.go` : creates a binary file in the same dir.
    - `./file` : to execute the binary from terminal

## Package : fmt

- `println()` : prints a line to the STDOUT, breaks the line after.
- `printf()` : prints the string and doesn't break the line at the end.
    - `printf("This line \n next line \n var = %v", var)`
    - Escape sequences : \n \v \t \\ \"

# Variables


## Types of declaration

### declaration with 'var' :

- scope : global
- initialize or declare

```go
var name string = "Sam" // declare
var address string // initialize

var name, address string = "Samuel","London" //declare multiple vars
```

### short variable declaration `:=`

- scope : local
- initializes

```go
name := "sam" //declare and initialize a variable

name,notName = "john","wane"
```

### When to use them

1. var :
    - bit operations
    - to declare and initialize
    - package level variables
    - when assigning functions to vars
2. := : 
    - inside a function
    - to initialize
    - Type doesn't matter
3. const :
    - cannot change the value once initialized
    - cannot declare

## Reassigning and Redeclaring

**allowed:**
```go
x := 10 //declaration
x = 30 //rassigning a value
```

**not allowed:**
```go
x := 10 //declaring
x := 20 //redeclaring another var with same name in the same scope
```

**allowed:**
```go
x := 10
x,y := 20,40 // Go sees this as : x=20 and y:=40
```


# Data types

## Integer

```go
var age int = -20 // - or +
bar age uint = 20 // + only
```
- Int types (bits of data the var can store):
    - *note* : since uint stores only +ve numbers, it can have twice as many +ve numbers as an int
    - int8 : (-128, 127)
    - uint8 : (0, 255)

---
## Float

```go
var price float32 = 3.5353 // less percise
var price float64 = 3.5353 // has more percision
```

---
## Strings

note: runes exist, just a remainder

```go
var myString string = "Hello \t World" //yes, excape sequences
var newString string = "Goodbye" + " " + "World"
```

get length: (number of bytes)
```go
fmt.Println(len("hello"))
```

---

## Booleans

```go
IAmAngry = true
IAmNotAngry = false
```

---
# Arithmetics

- You cannot add different datatypes together
- addition : +
- subtraction : -
- multiplication : *
- division : / (rounded down)
- modulus : %

```go
var age int = 34
var old_age float32 = 23.0
new_age = age+old_age // not allowed
new_age = float32(age)+old_age // conver a var to a common type
```

```go
var int1 = 5
var int2 = 3

result = int1/int2 // result is rounded down
```

