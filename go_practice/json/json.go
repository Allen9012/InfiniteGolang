package main

import (
	"encoding/json"
	"fmt"
)

// ADOrder struct
type ADOrderV2 struct {
	ADOrderID int64 `json:"adorder_id"`
	SvcID     int64 `json:"svc_id"`
}

// ADOrder struct
type ADOrderV1 struct {
	ADOrderID int64 `json:"adorder_id"`
}

func main() {
	adV1 := &ADOrderV1{
		ADOrderID: 1,
	}
	tmp := &ADOrderV2{}
	data, _ := json.Marshal(adV1)
	fmt.Println(string(data))
	err := json.Unmarshal(data, tmp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmp)
}
