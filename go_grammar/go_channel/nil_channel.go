package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan struct{}
	go func() {
		fmt.Println("hello1")
		<-c
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("hello2")
		c = make(chan struct{})
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("hello3")
		c <- struct{}{}
	}()

	time.Sleep(time.Second * 10)
}
