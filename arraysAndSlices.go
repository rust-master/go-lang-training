package main

import "fmt"

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