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
