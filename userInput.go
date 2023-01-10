package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() {
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	fmt.Println(name)

}