package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

/**
社区介绍控制器
*/
type IntroController struct {
	beego.Controller
}

/**
社区介绍页面
 */
func (c *IntroController) Get() {
	type Intro struct {
		Id     int    `json:"id"`
		Detail string `json:"detail"`
	}
	var intro Intro
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM ss_intro LIMIT 1").QueryRow(&intro)
	if err == nil {
		detail := strings.Replace(intro.Detail, "/file?", beego.AppConfig.String("file_path"), -1)
		intro.Detail = detail
		c.Data["data"] = intro
	}
	c.TplName = "intro/page.html"
}
