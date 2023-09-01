package data

import (
	"bytes"
	"fmt"
	"influxdb-test/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/influxdata/influxdb/client/v2"
)

func getVerifyData() {

}

const file_dir = "./data/"

func InitDb(clnt client.Client, db string, filename string) {

	q := client.NewQuery("create database "+db, "", "ns")
	if response, err := clnt.Query(q); err != nil || response.Error() != nil {
		beego.Error("create db error : ", err, response.Error())
	}
	url := common.HttpClnt.Addr + "/write?db=" + db
	method := "POST"
	dataFile := file_dir + filename
	// 读取数据文件内容
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Println("读取数据文件失败：", err)
		return
	}
	// 创建HTTP请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("创建HTTP请求失败:", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(common.Username, common.Password)
	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		beego.Error("发送HTTP请求失败:", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		// 请求不成功，处理请求失败
		beego.Error("write protocols 写入失败", resp.Status)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error("HTTP response reading failed: %s\n", err)
			return
		}
		beego.Error(fmt.Sprintf("HTTP response: %s\n", body))
		return
	}
	// 请求成功，处理返回结果
	beego.Informational("数据写入成功! 可支持write protocols")

}

// init test data
func Init() {
	clnt := common.HttpClnt.Client
	InitDb(clnt, "mydb", "test_data.txt")
	InitDb(clnt, "mydb", "NOAA_data.txt")
}
