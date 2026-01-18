package main

import (
	"fmt"
	"errors"
	"os"
)

func parentFunction() {
	fmt.Println("resulf of division : ", temp(4,5))
	
	c,d := returnT("string1", "string2")
	fmt.Println("returning two strings : ",c,d)

	// 4. blank identifier 
	result, _ := mulReturn(4.0,5.0) // ignore a value
	fmt.Println(result)
	
	deferOrder()
	deleteFiles("Filename", true)
	readFile()
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

// 5. Variadic functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 6. Defer

func readFile() {
	f,err := os.Open("Variables.md")
	if err != nil {
		return
	}
	defer f.Close() //executes when return gets called
	// defer runs after the return but finishes before function exits
}

// LIFO order (last deffered, first executed)
func deferOrder() {
	defer fmt.Println("Fourth")
	defer fmt.Println("Third")
	fmt.Println("First") // executes first
	defer fmt.Println("Second")
}


// 7. Panic (Rule of thump : DON't CALL PANIC unless necessary)

func deleteFiles(_ string, del bool) error {
	if del == true {
		//panic("You Can't delete the file")
		return errors.New("The file cannot be deleted")
	} else {
		return nil
	}
}
