package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
社区介绍控制器
*/
type IntroductionController struct {
	beego.Controller
}

/**
社区介绍页面
 */
func (c *IntroductionController) Get() {
	type Intro struct {
		Id     int    `json:"id"`
		Detail string `json:"detail"`
	}
	var intro Intro
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM ss_intro LIMIT 1").QueryRow(&intro)
	if err == nil {
		c.Data["data"] = intro
	}
	c.TplName = "introduction/page.html"
}
