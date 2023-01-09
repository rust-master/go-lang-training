package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hi, Go Lang")
	fmt.Println(quote.Go())

	// To Run:
	// go run main.go filename.go

	// variablesStringNums()
	// printingFormattingStrings()
	// arraysAndSlices()
	// standardLibrary()
	// loops()
	// booleansAndConditions()
	
	// var nums [5]int = [5]int{1,3,4,5,6}
	// findEvenFromArray(nums[:], printEven)

	// sum, sub := getMutlipleReturns(60, 30)
	// fmt.Printf("Sum is %v and sub is %v\n", sum, sub)
	
	// maps()
	// pointers()

	mybill := structAndCustomTypes("Zaryab")
	fmt.Println(mybill)
}

