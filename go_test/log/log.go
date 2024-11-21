package main

import (
	"fmt"
)

func main() {

	//ArcNotifyIcons := make(map[int64][]int32, 5)
	//
	//for i := 1; i < 5; i++ {
	//	ArcNotifyIcons[int64(i)] = []int32{1, 2, 4}
	//}
	//log.Fatalf("ArcNotifyIcons: %+v", ArcNotifyIcons)

	var amap map[int64][]int32

	if iconTypes, ok := amap[1234]; ok {
		fmt.Printf("iconTypes %v\n", iconTypes)
	}
}
