package main

import (
	"fmt"
	"errors"
)

func parentFunction() {
	fmt.Println("resulf of division : ", temp(4,5))
	
	c,d := returnT("string1", "string2")
	fmt.Println("returning two strings : ",c,d)

	// 4. blank identifier : ignores a return value
	result, _ := mulReturn()
}

// 1. naming multiple parameters consicutively
// specifying the return value's datatype
func temp(a,b int) int{ 
	return a+b
}

// 2. multiple return values
func returnT(c,d string) (string, string){
	return c,d
}

// 3. mutiple return statements
func mulReturn(e,f float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("Cannot Divide by Zero")
	}else {
		return e/f, nil
	}
}

