# Packages and Modules


```go
package main
func main() {}
```


## **Module** :

a collection of two or more packages

- `go mod init [directory]` : makes the dir a module
- `go.mod` : contains name of module version of go and external module

- `go run file.go` : creates a temporary binary file and executes the file, deleting the binary after execution
- `go build file.go` : creates a binary file in the same dir.
    - `./file` : to execute the binary from terminal


## **Package** : 

a named collection of Go files. Multiple packages in the same dir aren't allowed 

- `package main` : defines an executable program (has a main funct)
- `import "fmt"` : import built in packages 

import multiple packages : 

```go
import (
    "fmt"
    "strconv"
)
```

 
### package "fmt" #TODO

- `fmt.Println()` : prints a line to the STDOUT, breaks the line after.
- `fmt.Printf()` : prints the string and doesn't break the line at the end.
    - `printf("This line \n next line \n var = %v", var)`
    - Escape sequences : \n \v \t \\ \"
- `fmt.Scan(&var)` : Takes in string from the stdin and assigns the input to var
    - `fmt.Scanf(&var)` : takes in formatted string as input

### package "strconv"

- **output** : number, error saying whether the conversion was succesful
- `num,err := strconv.Atoi(str)`
