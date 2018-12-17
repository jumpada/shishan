package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
资讯控制器啊
*/
type InfoController struct {
	beego.Controller
}

/**
资讯页面
*/
func (c *InfoController) Page() {
	c.TplName = "info/page.html"
}

/**
资讯列表数据
*/
func (c *InfoController) List() {
	type Message struct {
		Id          int
		Title       string
		Cover       string
		Detail      string
		ReleaseTime string
	}
	var messages []Message
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail,
	release_time AS releaseTime
FROM
	message
WHERE
	category = 1
AND state = 1
ORDER BY
	release_time DESC
LIMIT ?,?`, (page-1)*limit, limit).QueryRows(&messages)
	if err == nil {
		c.Data["data"] = messages
	}
	c.ServeJSON()
}
