package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hi, Go Lang")
	fmt.Println(quote.Go())

	variablesStringNums()
}

func variablesStringNums() {

	var nameOne string = "M"
	var nameTwo = "Zaryab"
	var nameThree string // default its empty

	fmt.Println(nameOne, nameTwo, nameThree)
	nameThree = "Rafique"
	fmt.Println(nameOne, nameTwo, nameThree)

	// without the var keyword : shorthand for initializing and declaring the variable
	firstName := "Zaryab"
	fmt.Println(firstName)

	// ints 
	var age int = 20	
	var year = 2023
	month := 1
	fmt.Println(age, year, month)

	// bits and memory : https://pkg.go.dev/builtin@go1.19.4#IntegerType
	var numOne int8 = 127
	var numTwo int8 = -128
	fmt.Println(numOne, numTwo)

	// floats
	var width float32 = 25.5
	var height float64 = 100.1
	length := 2.4 // float64
	fmt.Println(width, height, length)
}