package controllers

import "github.com/astaxie/beego"

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
