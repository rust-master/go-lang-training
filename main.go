package main

import (
	"fmt"
	"sort"
	"strings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hi, Go Lang")
	fmt.Println(quote.Go())

	// variablesStringNums()
	// printingFormattingStrings()
	// arraysAndSlices()
	// standardLibrary()
	// loops()
	booleansAndConditions()
}

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

func loops() {
	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i<5; i++ {
		fmt.Println("value of i is:", i)
	}

	// loop throught slice
	coins := []string{"ETH", "BTC", "DOT", "ATOM", "SOL", "USDT", "USDC", "BNB", "AVAX"}

	for i :=0; i < len(coins); i++ {
		fmt.Println(coins[i])
	}

	// for in loop
	for index, value := range coins {
		fmt.Printf("%v at the position %v\n", value, index)
	}
}

func standardLibrary() {
	greeting := "Hello there Go Devs"

	fmt.Println(strings.Contains(greeting, "Hi"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello","Hi"))
	fmt.Println(strings.Index(greeting, "Go"))

	// sort
	age := []int{54,43,23,53,56,11,8,99,111}
	sort.Ints(age)
	fmt.Println(age)

	index := sort.SearchInts(age, 23)
	fmt.Println(index)
}

func arraysAndSlices() {
	var ages [7]int = [7]int{1,3,4}

	var years = [5]int{1,3,4}

	month := [6]string{"jan","feb","mar"}
	month[3] = "apr"

	fmt.Println(ages, len(ages))
	fmt.Println(years)
	fmt.Println(month)

	// slices (use arrays under the hood)
	var scores = []int{100, 40, 90}
	scores[1] = 30
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slices ranges
	rangeOne := scores[1:3] // not include the position 3
	fmt.Println(rangeOne, len(rangeOne))
 }

func printingFormattingStrings() {
	age := 20
	name := "Zaryab"
	// Print: no new line
	fmt.Print("Hello, ")
	fmt.Print("Go ")
	fmt.Print("Lang \n")

	// with new line
	fmt.Println("Hi Go") 
	fmt.Println(name, "age is", age)

	// Printf (formatted String) %_ = format specifeir
	fmt.Printf("my age is %v and my name is %v", age, name)

	// Sprintf (save formatted String)
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println("\nthe save string is:", str)
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