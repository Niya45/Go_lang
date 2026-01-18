package main

import (
	"fmt"
	"strconv"
	"time"
)

func parentControl() {
	//basic control structure:
	var age string
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	if numAge,errAge:=strconv.Atoi(age);errAge == nil {
		commentOnAge(numAge)
	}else {
		fmt.Println("Error : Invalid number")
	}
	
	//switch case:
	var day string
	fmt.Printf("What day is it?")
	fmt.Scan(&day)
	dayOfWeek(day)
}

// Baisic control structures
func commentOnAge(age int) {
	if age > 0 && age<18 {
		fmt.Println("You are young")
	}else if age>=18 && age<=30 {
		fmt.Println("You are an adult")
	}else if age>30 {
		fmt.Println("You are really old")
	}else {
		fmt.Println("You can't have that Age")
	}
}

// Logical Operators
func canDrive(age int, hasLicense bool) {
	if age>18 && hasLicense {
		fmt.Println("You can drive")
	}else if age<18 || !hasLicense {
		fmt.Println("You cannot drive")
	}
}

// Combining Operations
func canSlack(isWeekend, isHoliday bool, age int) {
	if (isWeekend || isHoliday) && age > 18 {
		fmt.Println("You can slack off")
	}
}


// Declare variable in statement (scope limited to if/else block)
func errCheck() {
	if _,err := strconv.Atoi("10"); err != nil {
		fmt.Println("Error:", err)
	}
}

// SWITCH case:
func dayOfWeek(day string) {
	switch day {
		// when day = case:
		case "Monday" :
			fmt.Println("First day of the week")
		case "Sunday","Saturday":
			fmt.Println("Weekend")
		// when all the cases are false:
		default:
			fmt.Println("You have Work/School right?")
	}
}

// Declare variable in switch
func dayOfWeekI() {
	switch day:=time.Now().Weekday(); day {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
	}
}

// Conditional Switch
func commentOnAgeSwitch(age int) {
	switch {
		case age < 18:
			fmt.Println("you are a child")
		case age > 18 && age < 30:
			fmt.Println("")
	}
}


