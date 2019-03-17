package main

import (
	_ "beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/go_video?charset=utf8")
	beego.Run()
}

