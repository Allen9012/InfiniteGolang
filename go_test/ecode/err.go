package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go-common/library/ecode"
	bm "go-common/library/net/http/blademaster"
	xtime "go-common/library/time"
	"time"
)

var (
	VideoupChargeMemberFirstCheckErr = ecode.Error(21554, "抢先看时间不可用，请修改后重试")
	VideoupStaffTypeNotExists        = ecode.Error(21088, "该分区暂未开放多人合作稿件")
	_missiontopic                    = "修改话题处于活动周期内，暂不支持修改，活动周期为%s"
	testErr                          = errors.New("这是一个err")
)

func main() {

	e := bm.NewServer(&bm.ServerConfig{
		Addr:    "0.0.0.0:8990",
		Timeout: xtime.Duration(time.Second),
	})

	err := errors.WithMessagef(VideoupChargeMemberFirstCheckErr, "抢先看时间不可用，请修改后重试")
	fmt.Printf("%+v\n", err)

	err2 := ecode.Errorf(ecode.RequestErr, "appid只能是数字")
	fmt.Printf("%+v\n", err2)

	err3 := errors.Wrapf(VideoupStaffTypeNotExists, "s.checkStaffType 分区%d不在配置里。配置：%+v", 123, "cache")

	err4 := errors.Wrapf(testErr, "s.checkStaffType 分区%d不在配置里。配置：%+v", 123, "cache")
	fmt.Printf("%+v\n", err3)

	fmt.Println(fmt.Sprintf(_missiontopic, time.Now().Format("2006年01月02日")))

	e.GET("/err1", func(c *bm.Context) {
		c.JSON(nil, err)
	})
	e.GET("/err2", func(c *bm.Context) {
		c.JSON(nil, err2)
	})
	e.GET("/err3", func(c *bm.Context) {
		c.JSON(nil, err3)
	})
	e.GET("/err4", func(c *bm.Context) {
		c.JSON(nil, err4)
	})
	if err := e.Start(); err != nil {
		panic(err)
	}
	select {}
}
