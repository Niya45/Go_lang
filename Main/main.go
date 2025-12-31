package main

import "fmt"

func main() {
	fmt.Println("This is a test run")
	// declare variables:
	var name string = "sam" // use in bit operations
	new_name := "jack" // to initialize a local variable
	// FMT package:
	fmt.Printf("My name used to be %v but now it's %v\n", name, new_name)
	// escape sequences:
	fmt.Printf("Line one \"backslash\"=\\ \nLine two \tTabspace \nLine three\v Vertical space\n")
}
