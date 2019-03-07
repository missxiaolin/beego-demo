package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id string
	Name string
}

func (this *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	id := this.GetString("name")
	fmt.Printf("%s\n", id)

	hash := map[string]interface{}{"success": 0,"message": "错误"}
	this.Data["json"] = hash
	this.ServeJSON()
}
