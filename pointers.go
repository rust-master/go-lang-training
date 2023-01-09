package main

import "fmt"

func pointers() {
	name := "Zaryab"
	n := &name // reference

	*n = "M Zaryab" // dereference

	fmt.Println("memory address", &name)
	fmt.Println("memory address", n)
	fmt.Println("value", *n)
}