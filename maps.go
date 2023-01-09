package main

import "fmt"

// maps are key-value
func maps() {

	coins_prices := map[string]int64{
		"BTC": 16000,
		"ETH": 1240,
		"SOL": 16,
		"DOT": 9,
	}

	fmt.Println(coins_prices)
	fmt.Println(coins_prices["ETH"])

	coins_prices["ETH"] = 1300
	coins_prices["AVAX"] = 20

	// loop through a map
	for k, v := range coins_prices {
		fmt.Println(k, "-", v)
	}

}