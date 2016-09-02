package common

import (
	"github.com/astaxie/beego"
)

type Error struct {
	Code    int32  `json:"code"`
	Msg     string `json:"msg"`
	CodeMsg string `json:"code_msg"`
}

type JsonResult struct {
	Code   int32  `json:"code"`
	Id     int64  `json:"id,omitempty"`
	ErrMsg string `json:"err_msg,omitempty"`
}

func Result(c beego.Controller, err error, id int64) {
	result := new(JsonResult)

	if err != nil {
		errResult := &JsonResult{1000, 0, "系统出错"}
		errResult.ErrMsg = err.Error()

		c.Data["json"] = errResult
		c.ServeJSON()
		return
	}

	result.Id = id

	c.Data["json"] = result
	c.ServeJSON()
	return
}

func InvalidResult(c beego.Controller, result *Error) {
	c.Data["json"] = result
	c.ServeJSON()
}
