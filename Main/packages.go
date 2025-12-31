// Packages and Modules:

package main 

import (
	"fmt"
	"strconv"
)

func fmtPack() {
	// FMT Package:
	fmt.Println("This is a line and I break at the end")
	var name string = "Sam"
	//formatted printing
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Backslash:\\ \nNewline Vertical\vspace\nTab\tspace \n")
}

func testVar() {
	// ceclaring vars : 
	var name string = "Sam"
	new_name := "Jack"

	fmt.Printf("Name : %v \nNew name : %v \n", name, new_name)
	
	//redeclaring and reassigning:
	age := 35
	//age := 36 (cannot redeclare a := variable)
	age, net_worth := 36,1000 
	fmt.Printf("Sam's age : %v \nSam's net worth : $%v\n", age, net_worth)
}

func arithmetic() {
	var inp1 string
	var inp2 string

	fmt.Printf("Enter first number : ")
	fmt.Scan(&inp1)
	fmt.Printf("Enter second number : ")
	fmt.Scan(&inp2)

	num1,err1 := strconv.Atoi(inp1)
	num2,err2 := strconv.Atoi(inp2)

	if err1 == nil && err2 == nil {
		fmt.Printf("adding : %v \n",num1+num2)
		fmt.Printf("subtracting : %v \n",num1-num2)
		fmt.Printf("multiplying : %v \n",num1*num2)
		fmt.Printf("dividing : %v R%v \n",num1/num2,num1%num2)
	}
}

