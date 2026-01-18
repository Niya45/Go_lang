// Packages and Modules:

package main 

import (
	"fmt"
	"strconv"
	"errors"
)

func parentFuncPackages() {
	fmtPack()
	//conv to int:
	var strToConv string
	fmt.Printf("Enter the \"number\" you wanna conver to int: ")
	fmt.Scan(&strToConv)
	convToInt(strToConv)

	//errors: 
		// if the error is only being used once, use := to declare and assign it together. 
		// Use var only if the error var is going to be reassigned
		// in a func(), return nil for a non-error outcome
	fmt.Println("TEST ERROR PACKAGE")
	var numer string
	var denomer string
	fmt.Printf("Enter numerator: ")
	fmt.Scan(&numer)
	fmt.Printf("Enter denominator: ")
	fmt.Scan(&denomer)

	numerInt,errTest1 := strconv.Atoi(numer)
	denomerInt,errTest2 := strconv.Atoi(denomer)

	if errTest1 == nil && errTest2 == nil {
		result,zeroErr := errorTest(numerInt, denomerInt)

		if zeroErr == nil {
			fmt.Print(result)
		}else {
			fmt.Print(zeroErr)
		}
	}
}

func fmtPack() {
	// FMT Package:
	fmt.Println("This is a line and I break at the end")
	var name string = "Sam"
	//formatted printing
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Backslash:\\ \nNewline Vertical\vspace\nTab\tspace \n")
}

func convToInt(strToConv string) {
	intResult,intErr := strconv.Atoi(strToConv)

	if intErr == nil{
		fmt.Printf("The result of converstion : %v ", intResult)
	}else{
		fmt.Print(intErr)
	}
}

func errorTest(numerator,denominator int) (int, error){
	if denominator == 0 {
		// creating a NEW type of error
		err := errors.New("Cannot Divide by Zero\n")
		return 0,err
	}else {
		return numerator/denominator, nil
	}
}

