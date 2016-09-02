package controllers

import (
	"idgo/common"
	"idgo/daos"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type IdController struct {
	beego.Controller
}

func (c *IdController) GetId() {
	data := new(daos.RedisDao)
	name := c.Ctx.Input.Param(":name")

	if name == "" || strings.Contains(name, " ") {
		common.InvalidResult(c.Controller, common.ValueError)
		return
	}

	id, err := data.GetIdByReids(name)

	common.Result(c.Controller, err, id)
}
