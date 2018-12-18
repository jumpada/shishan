package controllers

import (
	"github.com/astaxie/beego"
)

/**
狮山印象控制器
 */
type ImpressionController struct {
	beego.Controller
}

/**
狮山印象页面
 */
func (c *ImpressionController) Get() {
	c.TplName = "impression/page.html"
}
