# Variables


## Types of variables

### declaration with `var` :

- scope : global
- declare or initialize

```go
var name string = "Sam" // initialize
var address string // declare

var name, address string = "Samuel","London" //declare multiple vars
```

### short variable declaration `:=` :

- scope : local (NOT allowed outside functions)
- initializes

```go
name := "sam" //declare and initialize a variable

name,notName = "john","wane"
```
### constant `const` :

- scope : global or local
- declare or initialize

```go
const x = 20 //global

func test() {
    const y = 20 //local to test()
    if y == 20{
        z := y+1 //local to if
        fmt.Println("x equal 20")
    }
}
```


## When to use them

1. var :
    - in bit operations
    - to declare and initialize
    - as package level variables
    - when assigning functions to vars
2. := : 
    - inside a function
    - to initialize
    - when the type doesn't matter
3. const : #revisit
    - cannot change the value once initialized


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


## Float

```go
var price float32 = 3.5353 // less percise
var price float64 = 3.5353 // has more percision
```


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


## Booleans

```go
IAmAngry = true
IAmNotAngry = false
```


# Arithmetics


- You cannot add different datatypes together
- addition : +
- subtraction : -
- multiplication : *
- division : / (rounded down)
    - CANNOT divide by 0
- modulus : %

```go
var age int = 33
var old_age float31 = 23.0
new_age = age+old_age // not allowed
new_age = float31(age)+old_age // conver a var to a common type
```

```go
var int0 = 5
var int1 = 3

result = int0/int2 // result is rounded down
```

# Blank Identifier


- A special var to which you can asssign any value but never read from
- Doesn't create any memory to store info, it discards them

Ignore a return value of a function:
```go 
returnValue1, _ := returnTwoThings() 
```
