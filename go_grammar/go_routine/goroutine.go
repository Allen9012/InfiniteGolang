package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Goroutine")
	go func() {
		fmt.Println("Hello World")
		time.Sleep(time.Second * 2)
		fmt.Printf("time end")
	}()
	fmt.Println("main end")
}
