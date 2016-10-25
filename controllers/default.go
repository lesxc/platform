package controllers

import (
	"github.com/astaxie/beego"
)

// MainController  主要的控制层
type MainController struct {
	beego.Controller
}

// Get 返回多个值
func (c *MainController) Get() {
	c.TplName = "index.html"
}
