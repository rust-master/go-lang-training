package main

import "fmt"

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

