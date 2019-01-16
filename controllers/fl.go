package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type FlController struct {
	beego.Controller
}

func (c *FlController) Get() {
	type Fl struct {
		Id     int    `json:"id"`
		Detail string `json:"detail"`
	}
	var fl Fl
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM ss_fl LIMIT 1").QueryRow(&fl)
	if err == nil {
		c.Data["data"] = fl
	}
	c.TplName = "fl/page.html"
}
