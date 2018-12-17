package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "shishan/routers"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_url"), 30)
}
func main() {
	beego.Run()
}
