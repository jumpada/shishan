package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"time"
)

/**
咨询控制器
 */
type ConsultController struct {
	beego.Controller
}

/**
咨询页面
 */
func (c *ConsultController) Page() {
	openId := c.GetString("openId")
	if openId == "" {
		type Msg struct {
			Openid string
		}
		code := c.GetString("code")
		url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + beego.AppConfig.String("app_id") + "&secret=" + beego.AppConfig.String("app_secret") + "&code=" + code + "&grant_type=authorization_code"
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		if bytes.Contains(body, []byte("errcode")) {
			return
		} else {
			msg := Msg{}
			err = json.Unmarshal(body, &msg)
			if err == nil {
				c.Data["openId"] = msg.Openid
			}
		}
	} else {
		c.Data["openId"] = openId
	}
	c.TplName = "consult/page.html"
}

/**
咨询表单
 */
func (c *ConsultController) Form() {
	openId := c.GetString("openId")
	c.Data["openId"] = openId;
	c.TplName = "consult/form.html"
}

/**
咨询表单提交
 */
func (c *ConsultController) Submit() {
	title := c.GetString("title")
	content := c.GetString("content")
	contact := c.GetString("contact")
	openId := c.GetString("openId")
	o := orm.NewOrm()
	res, err := o.Raw(`INSERT INTO ss_consult (
	title,
	content,
	open_id,
	contact,
	consult_time
)
VALUES
	(?,?,?,?,?)`, title, content, openId, contact, time.Now()).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		c.Data["json"] = map[string]interface{}{"state": num == 1}
	} else {
		c.Data["json"] = map[string]interface{}{"state": false}
	}
	c.ServeJSON()
}

/**
咨询历史记录
 */
func (c *ConsultController) History() {
	openId := c.GetString("openId")
	c.Data["openId"] = openId;
	c.TplName = "consult/history.html"
}

/**
列表数据
 */
func (c *ConsultController) List() {
	type Consult struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Content     string `json:"content"`
		State       string `json:"state"`
		ConsultTime string `json:"consultTime"`
	}
	var consultArr []Consult
	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")
	openId := c.GetString("openId")
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	content,
	state,
	DATE_FORMAT(
		consult_time,
		'%Y-%m-%d %H:%i:%s'
	) AS consult_time
FROM
	ss_consult
WHERE
	open_id =?
ORDER BY
	consult_time DESC
LIMIT ?,?`, openId, (page-1)*pageSize, pageSize).QueryRows(&consultArr)
	if err == nil {
		c.Data["json"] = consultArr
	}
	c.ServeJSON()
}

/**
咨询详细
 */
func (c *ConsultController) Detail() {
	type Consult struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Content     string `json:"content"`
		State       string `json:"state"`
		ConsultTime string `json:"consultTime"`
		Reply       string `json:"reply"`
		ReplyTime   string `json:"replyTime"`
	}
	var consultArr Consult
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	title,
	content,
	state,
	DATE_FORMAT(
		consult_time,
		'%Y-%m-%d %H:%i:%s'
	) AS consult_time,
	reply,
	DATE_FORMAT(
		reply_time,
		'%Y-%m-%d %H:%i:%s'
	) AS reply_time
FROM
	ss_consult
WHERE
	id =?`, id).QueryRow(&consultArr)
	if err == nil {
		c.Data["data"] = consultArr
	}
	c.TplName = "consult/detail.html"
}
