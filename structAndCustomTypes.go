package main

// struct (custom types)
type bill struct {
	name string
	items map[string]float64
	tip float64
}

func structAndCustomTypes(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}