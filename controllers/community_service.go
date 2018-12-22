package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
社区动态控制器
 */
type CommunityServiceController struct {
	beego.Controller
}

/**
社区动态页面
 */
func (c *CommunityServiceController) Page() {
	c.TplName = "community_service/page.html"
}

/**
社区动态数据
 */
func (c *CommunityServiceController) List() {
	type CommunityService struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		Detail      string `json:"detail"`
		ReleaseTime string `json:"releaseTime"`
	}
	var csArr []CommunityService
	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")
	category, _ := c.GetInt("category")
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail,
	DATE_FORMAT(release_time, '%Y-%m-%d %H:%i:%s') AS release_time
FROM
	ss_community_service
WHERE
	category = ?
AND state = 1
ORDER BY
	release_time DESC
LIMIT ?,?`, category, (page-1)*pageSize, pageSize).QueryRows(&csArr)
	if err == nil {
		c.Data["json"] = csArr
	}
	c.ServeJSON()
}

/**
详细页面
 */
func (c *CommunityServiceController) Detail() {
	id := c.GetString("id")
	type CommunityService struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		Detail      string `json:"detail"`
		ReleaseTime string `json:"releaseTime"`
	}
	var communityService CommunityService
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	title,
	cover,
	detail,
	DATE_FORMAT(release_time, '%Y-%m-%d %H:%i:%s') AS release_time
FROM
	ss_community_service WHERE id=?`, id).QueryRow(&communityService)
	if err == nil {
		c.Data["data"] = communityService
	}
	c.TplName = "community_service/detail.html"
}
