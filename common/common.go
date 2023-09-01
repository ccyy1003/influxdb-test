package common

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	Addr     string
	Username string
	Password string
	Version  string
	Token    string
)

func CheckEnv() {

	Addr = os.Getenv("INFLUX_TEST_ADDR")
	if Addr == "" {
		Addr = "http://127.0.0.1:8086"
	}
	Username = os.Getenv("INFLUX_TEST_USER")
	Password = os.Getenv("INFLUX_TEST_PWD")
	if Username != "" && Password != "" {
		Token = Username + ":" + Password
	} else {
		Token = ""
	}
}
func SetLog(name string) {
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile,
		`{"filename":"/var/log/influxdb_test/`+name+`.log","level":7,"daily":true,"maxdays":7}`)
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true
	beego.BConfig.Listen.ServerTimeOut = 60
	beego.BConfig.Listen.ListenTCP4 = true
}
