package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

/**
社区概况控制器
*/
type OverviewController struct {
	beego.Controller
}

func (c *OverviewController) Page() {
	type Overview struct {
		Id     int
		Detail string
	}
	var overview Overview
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM overview LIMIT 1").QueryRow(&overview)
	if err == nil {
		detail := strings.Replace(overview.Detail, "/file?", beego.AppConfig.String("file_path"), -1)
		overview.Detail = detail
		c.Data["data"] = overview
	}
	c.TplName = "overview/page.html"
}
