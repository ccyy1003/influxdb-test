package controllers

import (
	"bytes"
	"encoding/json"
	"influxdb-test/common"
	"net/http"
	"os"

	"github.com/astaxie/beego"
)

// BaseAPIView ...
type BaseAPIView struct {
	beego.Controller
	tr            interface{}
	noRequireAuth bool   // 不需要验证
	apidatainit   bool   // 是否已初始化
	format        string // 请求数据返回格式 json or api
	body          []byte // out body
}

var viaHost string

func init() {
	if h, err := os.Hostname(); err == nil {
		viaHost = h
	}
}

// noNeedAuth 不需要认证
func (o *BaseAPIView) noNeedAuth() {
	o.noRequireAuth = true
}

// getformat 获取请求格式
func (o *BaseAPIView) getformat() {
	o.format = "json"
}

// initRequest ...
func (o *BaseAPIView) initRequest() error {

	// get format
	o.getformat()

	// Server
	s := beego.AppConfig.String("appname")
	o.Ctx.Output.Header("Via", "http/1.1 "+viaHost+" ("+s+" API Server/"+common.Version+")")
	o.Ctx.Output.Header("Server", s+"/"+common.Version)

	o.apidatainit = true

	return nil
}

// ServeJSONMixin ...
func (o *BaseAPIView) ServeJSONMixin() {
	o.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	o.Ctx.Output.Body(o.body)
}

// RespRedirect ...
func (o *BaseAPIView) RespRedirect(location string, data interface{}, decode ...bool) {
	o.Ctx.Output.SetStatus(302)
	o.Ctx.Output.Header("Location", location)
	o.respbody(0, "", data, decode)
	o.ServeJSONMixin()
}

// RespOK ...
func (o *BaseAPIView) RespOK(data interface{}, decode ...bool) {
	o.Ctx.Output.SetStatus(200)
	o.respbody(0, "", data, decode)
	o.ServeJSONMixin()

}

// RespError 返回错误
func (o *BaseAPIView) RespError(code int, message string, decode ...bool) {
	o.Ctx.Output.Status = code
	o.respbody(code, message, nil, decode)
	o.ServeJSONMixin()
}

func (o *BaseAPIView) respbody(code int, message string, data interface{}, decode []bool) {
	if !o.apidatainit {
		panic("Controller api data need init")
	}
	if data != nil {
		if temp, ok := data.(common.TestRes); ok {
			o.tr = temp
		} else if temp, ok := data.([]common.TestRes); ok {
			o.tr = temp
		}
	}

	o.Data["json"] = o.tr

	if len(decode) == 1 && decode[0] {
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		if beego.BConfig.RunMode == "dev" {
			jsonEncoder.SetIndent("", " ")
		}
		jsonEncoder.Encode(o.Data["json"])
		o.body = bf.Bytes()
	} else {
		var body []byte
		var err error
		if beego.BConfig.RunMode == "prod" {
			body, err = json.MarshalIndent(o.Data["json"], "", "  ")
			//body, err = json.Marshal(o.Data["json"])
		} else {
			body, err = json.MarshalIndent(o.Data["json"], "", "  ")
		}
		if err != nil {
			http.Error(o.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		o.body = body
	}
}
