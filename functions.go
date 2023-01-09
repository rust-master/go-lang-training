package main

import "fmt"

func getMutlipleReturns(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}

func printEven(){
	fmt.Println("even")
}

func findEvenFromArray(nums []int, f func()) {
	for i :=0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Printf("%v is the ", nums[i])
			f()
		}
	}
}
