package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

// Get 返回多个值
func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	var user string
	c.Data["user"]= user
	fmt.Println(user)
	c.TplName = "index.html"


}
