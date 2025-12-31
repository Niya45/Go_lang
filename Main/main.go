package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Basic stuff:")
	// declare variables:
	var name string = "sam" // use in bit operations
	new_name := "jack" // to initialize a local variable
	// FMT package:
	fmt.Printf("My name used to be %v but now it's %v\n", name, new_name)
	// escape sequences:
	fmt.Printf("Line one \"backslash\"=\\ \nLine two \tTabspace \nLine three\v Vertical space\n")
	fmt.Printf("==END BASIC STUFF===\n\n")


	var inp1 string
	var inp2 string

	fmt.Println("Enter first number :")
	fmt.Scan(&inp1)
	
	fmt.Println("Enter second number :")
	fmt.Scan(&inp2)

	// arithmetic function:
	arithmetic(inp1, inp2)
}

func arithmetic(inp1, inp2 string) {

	num1,err1 := strconv.Atoi(inp1)	
	num2,err2 := strconv.Atoi(inp2)
	//addition:
	sum := num1+num2
	//subtraction:
	min := num2-num1
	//multiplication:
	product := num1*num2
	//division:
	div := num2/num1
	remi := num2%num1

	// if-conditional:
	if err1 == nil && err2 == nil {
		fmt.Printf("The sum of numbers : %v \n", sum)
		fmt.Printf("The difference of numbers : %v \n", min)
		fmt.Printf("The product of numbers : %v \n", product)
		fmt.Printf("The result of divison : %vR%v \n", div, remi)
	}
}
