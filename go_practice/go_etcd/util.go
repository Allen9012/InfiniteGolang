package main

import (
	"fmt"
	"strconv"
)

func fromUint64(number uint64) string {
	return strconv.Itoa(int(number))
}

func StoUint64(value string) uint64 {
	ret, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("value to toUint64 failed, err: %v\n", err)
		return 0
	}
	return uint64(ret)
}

func BtoUInt64(value []byte) uint64 {
	ret, err := strconv.Atoi(string(value))
	if err != nil {
		fmt.Printf("value to uint64 failed, err: %v\n", err)
		return 0
	}
	return uint64(ret)
}
