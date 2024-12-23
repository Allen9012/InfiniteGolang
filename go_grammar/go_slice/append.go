package main

import "fmt"

func main() {

	//x := make([]int64, 0, 4)
	var x []int64

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y := append(x, 1)
	z := append(x, 2)
	fmt.Printf("x:%+v,%p cap:%d, len:%d \n", x, x, cap(x), len(x))
	fmt.Printf("y:%+v,%p cap:%d, len:%d\n", y, y, cap(y), len(y))
	fmt.Printf("z:%+v,%p cap:%d, len:%d \n", z, z, cap(z), len(z))
}
