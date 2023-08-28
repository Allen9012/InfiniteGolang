package main

import (
	"errors"
	"fmt"
)

// 使用%w 来表示包装一个e
func main() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误: %w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
}
