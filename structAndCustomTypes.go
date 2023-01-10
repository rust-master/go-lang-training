package main

import "fmt"

// struct (custom types)
type bill struct {
	name string
	items map[string]float64
	tip float64
}

// Receiver Functions
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ....%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ....$%0.2f\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ....$%0.2f", "total:", total + b.tip)

	return fs
}

func structAndCustomTypes(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}

	return b
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip  // automatice dereference by go lang
}

// Receiver Functions with Pointers
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}