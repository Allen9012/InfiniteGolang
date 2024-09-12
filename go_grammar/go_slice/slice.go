package main

import "fmt"

func returnSlice() (s []int) {
	return s
}

func main() {
	// 测试切片
	slice := returnSlice()
	fmt.Println("Slice length:", len(slice)) // 如果切片为 nil，返回 0
	if slice == nil {
		fmt.Println("Slice is nil")
	}
}
