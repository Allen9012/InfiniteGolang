package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Duration(1) * time.Millisecond
	fmt.Println(duration) // 1ms
	duration2 := time.Duration(1)
	fmt.Println(duration2) // 1ns
}
