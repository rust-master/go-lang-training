package main

import (
	"fmt"
	"sort"
	"strings"
)

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