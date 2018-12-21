package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
办事指南控制器
 */
type GuidanceController struct {
	beego.Controller
}

/**
办事指南页面
 */
func (c *GuidanceController) Page() {
	type Guidance struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Cover  string `json:"cover"`
		Detail string `json:"detail"`
	}
	var guidanceArr []Guidance
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail
FROM
	ss_guidance
ORDER BY
	create_time DESC`).QueryRows(&guidanceArr)
	if err == nil {
		c.Data["data"] = guidanceArr
	}
	c.TplName = "guidance/page.html"
}

/**
详细页面
 */
func (c *GuidanceController) Detail() {
	id := c.GetString("id")
	type Guidance struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Cover  string `json:"cover"`
		Detail string `json:"detail"`
	}
	var guidance Guidance
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail
FROM
	ss_guidance WHERE id=?`, id).QueryRow(&guidance)
	if err == nil {
		c.Data["data"] = guidance
	}
	c.TplName = "guidance/detail.html"
}
