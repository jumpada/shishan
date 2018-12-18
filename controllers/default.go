package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

/**
默认控制器
*/
type DefaultController struct {
	beego.Controller
}

/**
微信接入认证
*/
func (c *DefaultController) Get() {
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	signature := c.GetString("signature")
	echostr := c.GetString("echostr")
	arr := []string{beego.AppConfig.String("wechat_token"), timestamp, nonce}
	sort.Strings(arr)
	arrStr := strings.Join(arr, "")
	s := sha1.New()
	_, _ = io.WriteString(s, arrStr)
	signatureGen := fmt.Sprintf("%x", s.Sum(nil))
	if signatureGen == signature {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("access failed")
	}
}

/**
微信消息回复
 */
func (c *DefaultController) Post() {
	c.Ctx.WriteString("ok")
}
