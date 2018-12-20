package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

/**
文件控制器
 */
type FileController struct {
	beego.Controller
}

/**
文件下载
 */
func (c *FileController) Get() {
	filePath := c.GetString("filePath")
	rs := []rune(filePath)
	startIndex := strings.LastIndex(filePath, "/")
	endIndex := len(filePath)
	c.Ctx.Output.Download(beego.AppConfig.String("file_path")+filePath, string(rs[startIndex:endIndex]))
}
