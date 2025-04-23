package main

import (
	"fmt"
	bm "go-common/library/net/http/blademaster"
	xtime "go-common/library/time"
	"time"
)

type User struct {
	Name string `form:"name" json:"name"`
	Uid  int64  `form:"uid" json:"uid"`
}

type UserValidate struct {
	Name string `form:"name" json:"name" validate:"required"`
	Uid  int64  `form:"uid" json:"uid" validate:"min=0,max=100"`
}

func Init() (engine *bm.Engine) {
	e := bm.NewServer(&bm.ServerConfig{
		Addr:    "0.0.0.0:8990",
		Timeout: xtime.Duration(time.Second),
	})
	e.GET("/notvalidate", func(c *bm.Context) {
		data := &User{}
		err := c.Bind(data)
		if err != nil {
			c.JSON(nil, err)
			fmt.Println(err)
		}
		fmt.Println(data)
	})
	e.GET("/validate", func(c *bm.Context) {
		data := &UserValidate{}
		err := c.Bind(data)
		if err != nil {
			c.JSON(nil, err)
			fmt.Println(err)
		}
		fmt.Println(data)
	})
	if err := e.Start(); err != nil {
		panic(err)
	}
	select {}
}

//  curl http://0.0.0.0:8990/validate\?name\=testuser\&uid\=123456

func main() {
	Init()
}
