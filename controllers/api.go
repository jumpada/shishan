package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type ApiController struct {
	beego.Controller
}

/**
登录
 */
func (c *ApiController) Login() {
	result := make(map[string]interface{})
	username := c.GetString("username")
	password := c.GetString("password")
	type User struct {
		Username string
		Password string
	}
	var user User
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	username,
	original_pwd AS password
FROM
	base_user
WHERE
	username =?`, username).QueryRow(&user)
	if err == nil {
		if user.Password == password {
			result["state"] = true
		} else {
			result["state"] = false
			result["code"] = 2
		}
	} else {
		result["state"] = false
		result["code"] = 1

	}
	c.Data["json"] = result
	c.ServeJSON()
}

/**
志愿者活动列表
 */
func (c *ApiController) VoluntaryList() {
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
	ss_voluntary
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
志愿者信息
 */
func (c *ApiController) VolunteerInfo() {
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
	ss_volunteer
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
志愿者活动签到
 */
func (c *ApiController) VoluntarySign() {
	volunteerId, _ := c.GetInt("volunteerId")
	voluntaryId, _ := c.GetInt("voluntaryId")
	o := orm.NewOrm()
	_ = o.Begin()
	type Count struct {
		Num int
	}
	type Voluntary struct {
		Score int
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
						var voluntary Voluntary
						err := o.Raw(`SELECT
									score
									FROM
										ss_voluntary
									WHERE
									id =?`, voluntaryId).QueryRow(&voluntary)
						if err == nil {
							res, err = o.Raw(`UPDATE ss_volunteer
											SET score = score + ?
											WHERE
											id = ?`, voluntary.Score, volunteerId).Exec()
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
