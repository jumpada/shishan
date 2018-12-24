package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
志愿者活动控制器
 */
type VoluntaryController struct {
	beego.Controller
}

func (c *VoluntaryController) Page() {
	c.TplName = "voluntary/page.html"
}

func (c *VoluntaryController) List() {
	type Voluntary struct {
		Id             int    `json:"id"`
		Name           string `json:"name"`
		Target         string `json:"target"`
		Address        string `json:"address"`
		Num            int    `json:"num"`
		TimeLimit      string `json:"timeLimit"`
		TimeLimitCount string `json:"timeLimitCount"`
		Detail         string `json:"detail"`
		State          int    `json:"state"`
		CreateTime     string `json:"createTime"`
	}
	var voluntaryArr []Voluntary
	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")
	state, _ := c.GetInt("state")
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	name,
	target,
	address,
	num,
	DATE_FORMAT(
		time_limit,
		'%Y-%m-%d %H:%i:%s'
	) AS time_limit,
	DATE_FORMAT(
		time_limit,
		'%m/%d/%Y %H:%i:%s'
	) AS time_limit_count,
	detail,
	state,
	DATE_FORMAT(
		create_time,
		'%Y-%m-%d %H:%i:%s'
	) AS create_time
FROM
	voluntary
WHERE
	state = ?
ORDER BY
	create_time DESC
LIMIT ?,?`, state, (page-1)*pageSize, pageSize).QueryRows(&voluntaryArr)
	if err == nil {
		c.Data["json"] = voluntaryArr
	}
	c.ServeJSON()
}

/**
详细页面
 */
func (c *VoluntaryController) Detail() {
	id := c.GetString("id")
	type Voluntary struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Target     string `json:"target"`
		Address    string `json:"address"`
		Num        int    `json:"num"`
		TimeLimit  string `json:"timeLimit"`
		Detail     string `json:"detail"`
		State      int    `json:"state"`
		CreateTime string `json:"createTime"`
	}
	var voluntary Voluntary
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	name,
	target,
	address,
	num,
	DATE_FORMAT(
		time_limit,
		'%Y-%m-%d %H:%i:%s'
	) AS time_limit,
	detail,
	state,
	DATE_FORMAT(
		create_time,
		'%Y-%m-%d %H:%i:%s'
	) AS create_time
FROM
	voluntary
WHERE
	id =?`, id).QueryRow(&voluntary)
	if err == nil {
		c.Data["data"] = voluntary
	}
	c.TplName = "voluntary/detail.html"
}
