package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
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
	type Msg struct {
		OpenId string
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
		if err != nil {
			c.Data["open_id"] = msg.OpenId
		}
	}
	c.TplName = "consult/page.html"
}
