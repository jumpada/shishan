package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
党在身边控制器
 */
type PartySideController struct {
	beego.Controller
}

/**
党在身边页面
 */
func (c *PartySideController) Page() {
	type PartySide struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		Detail      string `json:"detail"`
		ReleaseTime string `json:"releaseTime"`
	}
	var partySideArr []PartySide
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail,
	DATE_FORMAT(release_time, '%Y-%m-%d %H:%i:%s') AS release_time
FROM
	ss_party_side
WHERE state = 1
ORDER BY
	create_time DESC`).QueryRows(&partySideArr)
	if err == nil {
		c.Data["data"] = partySideArr
	}
	c.TplName = "party_side/page.html"
}

/**
详细页面
 */
func (c *PartySideController) Detail() {
	id := c.GetString("id")
	type PartySide struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		Detail      string `json:"detail"`
		ReleaseTime string `json:"releaseTime"`
	}
	var partySide PartySide
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail,
	DATE_FORMAT(release_time, '%Y-%m-%d %H:%i:%s') AS release_time
FROM
	ss_party_side WHERE id=?`, id).QueryRow(&partySide)
	if err == nil {
		c.Data["data"] = partySide
	}
	c.TplName = "party_side/detail.html"
}
