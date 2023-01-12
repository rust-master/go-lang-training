package main

import (
	"fmt"
	"time"
)

func doSomething(size int) {
    for i := 0; i < size; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}