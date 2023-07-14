package main

import (
	"fmt"
	"time"
)

// 通道上的range循环：接收通道中的值的常见习惯是使用for range循环。当使用for range循环遍历通道时，它会自动检测通道是否已关闭。一旦通道关闭并且所有值都已接收，循环将退出。
func main() {

	ch := make(chan int) //	创建一个channel
	go receive(ch)       //	开启一个goroutine去接收channel的数据
	go sender(ch)        //	开启一个goroutine去发送数据

	select {
	case <-time.After(time.Second * 10):
		println("timeout")
	}
}

func receive(ch <-chan int) {
	for i := range ch {
		println("Received:", i)
	}
	fmt.Println("range done")
}

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}
