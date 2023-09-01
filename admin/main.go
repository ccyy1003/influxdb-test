package main

import (
	"fmt"
	"influxdb-test/common"
	_ "influxdb-test/routers"

	"github.com/astaxie/beego"
)

var HttpClnt = common.HttpClnt

func main() {
	common.SetLog("admin")
	defer func() {
		if errs := recover(); errs != nil {
			fmt.Println(errs)
		}
	}()

	// env check
	common.CheckEnv()
	if HttpClnt.Init() != nil {
		beego.Error("httpclnt init error")
		return
	}
	defer HttpClnt.Client.Close()

	beego.Run()

}
