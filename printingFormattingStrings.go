package main

import "fmt"

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