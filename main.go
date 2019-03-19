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
	beego.SetStaticPath("/resources", "resources")
	beego.SetStaticPath("/.well-known", ".well-known")
	beego.SetStaticPath("/MP_verify_1oIq6VPv7tFQDiyp.txt", "MP_verify_1oIq6VPv7tFQDiyp.txt")
	beego.Run()
}
