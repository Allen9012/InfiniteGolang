package main

import "fmt"

func main() {

	var rawMap map[string]string

	if _, ok := rawMap["test"]; !ok {
		fmt.Println("ok")
	}
	fmt.Println(rawMap)
}
