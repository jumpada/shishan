package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DwController struct {
	beego.Controller
}

func (c *DwController) Get() {
	type Dw struct {
		Id     int    `json:"id"`
		Detail string `json:"detail"`
	}
	var dw Dw
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM ss_dw LIMIT 1").QueryRow(&dw)
	if err == nil {
		c.Data["data"] = dw
	}
	c.TplName = "dw/page.html"
}
