package main

import "fmt"

func booleansAndConditions() {
	age := 20
	
	fmt.Println(age <= 15)
	fmt.Println(age >= 10)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is ", age)
	}
}