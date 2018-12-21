package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
党员活动控制器
 */
type PartyActivityController struct {
	beego.Controller
}

/**
党员活动页面
 */
func (c *PartyActivityController) Page() {
	type PartyActivity struct {
		Id           int    `json:"id"`
		Title        string `json:"title"`
		ActivityTime string `json:"activityTime"`
		ActivitySite string `json:"activitySite"`
		Attendee     string `json:"attendee"`
		Principal    string `json:"principal"`
		Phone        string `json:"phone"`
		Remark       string `json:"remark"`
	}
	var partyActivityArr []PartyActivity
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	activity_time,
	activity_site,
	attendee,
	principal,
	phone,
	remark
FROM
	ss_party_activity
ORDER BY
	create_time DESC`).QueryRows(&partyActivityArr)
	if err == nil {
		c.Data["data"] = partyActivityArr
	}
	c.TplName = "party_activity/page.html"
}
