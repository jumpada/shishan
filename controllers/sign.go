package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

/**
签到控制器
 */
type SignController struct {
	beego.Controller
}

/**
志愿者信息
 */
func (c *SignController) VolunteerInfo() {
	result := make(map[string]interface{})
	id, err := c.GetInt("id")
	if err == nil {
		type Volunteer struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			IdNumber string `json:"idNumber"`
		}
		var volunteer Volunteer
		o := orm.NewOrm()
		err := o.Raw(`SELECT
	id,
	name,
	id_number
FROM
	volunteer
WHERE
	id =?`, id).QueryRow(&volunteer)
		if err == nil {
			result["state"] = true
			result["info"] = volunteer
		} else {
			result["state"] = false
		}
	} else {
		result["state"] = false
	}
	c.Data["json"] = result
	c.ServeJSON()
}

/**
志愿者活动列表
 */
func (c *SignController) VoluntaryList() {
	type Voluntary struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var voluntaryArr []Voluntary
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	name
FROM
	voluntary
WHERE
	state = 1
ORDER BY
	create_time DESC`).QueryRows(&voluntaryArr)
	if err == nil {
		c.Data["json"] = voluntaryArr
	}
	c.ServeJSON()
}

/**
签到
 */
func (c *SignController) Sign() {
	volunteerId, _ := c.GetInt("volunteerId")
	voluntaryId, _ := c.GetInt("voluntaryId")
	o := orm.NewOrm()
	_ = o.Begin()
	type Count struct {
		Num int
	}
	var count Count
	err := o.Raw(`SELECT
	COUNT(*) AS num
FROM
	ss_sign
WHERE
	volunteer_id =?
AND voluntary_id =?`, volunteerId, voluntaryId).QueryRow(&count)
	if err == nil {
		if count.Num > 0 {
			c.Data["json"] = map[string]interface{}{"state": false, "code": 1}
			_ = o.Rollback()
		} else {
			res, err := o.Raw(`INSERT INTO ss_sign (
	volunteer_id,
	voluntary_id,
	sign_time
)
VALUES
	(?,?,?)`, volunteerId, voluntaryId, time.Now()).Exec()
			if err == nil {
				num, err := res.RowsAffected()
				if err == nil {
					if num == 1 {
						res, err = o.Raw(`UPDATE volunteer
SET score = score + 1
WHERE
	id = ?`, volunteerId).Exec()
						if err == nil {
							num, err := res.RowsAffected()
							if err == nil {
								if num == 1 {
									c.Data["json"] = map[string]interface{}{"state": true}
									_ = o.Commit()
								} else {
									c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
									_ = o.Rollback()
								}
							} else {
								c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
								_ = o.Rollback()
							}
						} else {
							c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
							_ = o.Rollback()
						}
					} else {
						c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
						_ = o.Rollback()
					}
				} else {
					c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
					_ = o.Rollback()
				}
			} else {
				c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
				_ = o.Rollback()
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"state": false, "code": 2}
		_ = o.Rollback()
	}
	c.ServeJSON()
}
