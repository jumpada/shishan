package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"strings"
)

type VoteController struct {
	beego.Controller
}

func (c *VoteController) Page() {
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
	c.TplName = "vote/page.html"
}

func (c *VoteController) List() {
	type Vote struct {
		Id             string `json:"id"`
		Title          string `json:"title"`
		Description    string `json:"description"`
		TimeLimitCount string `json:"timeLimitCount"`
	}
	var voteArr []Vote
	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")
	o := orm.NewOrm()
	_, err := o.Raw(`SELECT
	id,
	title,
	description,
	DATE_FORMAT(
		time_limit,
		'%m/%d/%Y %H:%i:%s'
	) AS time_limit_count
FROM
	vote
WHERE
	state = 1
ORDER BY
	release_time DESC
LIMIT ?,?`, (page-1)*pageSize, pageSize).QueryRows(&voteArr)
	if err == nil {
		c.Data["json"] = voteArr
	}
	c.ServeJSON()
}

func (c *VoteController) Detail() {
	id := c.GetString("id")
	openId := c.GetString("openId")
	c.Data["openId"] = openId
	type Vote struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		TimeLimit   string `json:"timeLimit"`
		Category    int    `json:"category"`
		State       int    `json:"state"`
	}
	type VoteOption struct {
		Id       int    `json:"id"`
		Option   string `json:"option"`
		Selected bool   `json:"selected"`
	}
	type Voted struct {
		Num int `json:"num"`
	}
	var vote Vote
	o := orm.NewOrm()
	err := o.Raw(`SELECT
	id,
	title,
	description,
	DATE_FORMAT(
		time_limit,
		'%Y-%m-%d %H:%i:%s'
	) AS time_limit,
	category,
	state
FROM
	vote
WHERE
	id =?`, id).QueryRow(&vote)
	if err == nil {
		c.Data["vote"] = vote
	}
	var options []VoteOption
	var voted Voted
	_, err1 := o.Raw("SELECT id,`option` FROM vote_option WHERE vote_id=?", id).QueryRows(&options)
	if err1 == nil {
		var ops []VoteOption
		for _, option := range options {
			op := VoteOption{}
			op.Id = option.Id
			op.Option = option.Option
			_ = o.Raw(`SELECT
	count(*) AS num
FROM
	voted
WHERE
	open_id =?
AND option_id =?`, openId, option.Id).QueryRow(&voted)
			op.Selected = voted.Num > 0
			ops = append(ops, op)
		}
		c.Data["options"] = ops
	}
	err2 := o.Raw(`SELECT
	count(*) AS num
FROM
	voted A
LEFT JOIN vote_option B ON A.option_id = B.id
LEFT JOIN vote C ON B.vote_id = C.id
WHERE
	A.open_id =?
AND C.id =?`, openId, id).QueryRow(&voted)
	if err2 == nil {
		c.Data["voted"] = voted.Num > 0
	}
	c.TplName = "vote/detail.html"
}

func (c *VoteController) Submit() {
	options := c.GetString("options")
	openId := c.GetString("openId")
	optionArr := strings.Split(options, "☻")
	o := orm.NewOrm()
	for _, value := range optionArr {
		_, _ = o.Raw(`INSERT INTO voted (open_id, option_id)
VALUES
	(?,?)`, openId, value).Exec()
	}
	c.Data["json"] = map[string]interface{}{"state": true}
	c.ServeJSON()
}
