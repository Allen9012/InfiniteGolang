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

	var a = []int64{1, 2, 3, 4}
	var b []int64
	a = append(a, b...)
	//b = append(b, a...)
	fmt.Printf("a:%+v,%p cap:%d, len:%d \n", a, a, cap(a), len(a))
	fmt.Printf("b:%+v,%p cap:%d, len:%d \n", b, b, cap(b), len(b))
	var c []int64 = nil
	fmt.Printf("len(%d)\n", len(c))
}
