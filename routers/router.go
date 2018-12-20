package routers

import (
	"github.com/astaxie/beego"
	"shishan/controllers"
)

func init() {
	//微信接入验证
	beego.Router("/access", &controllers.DefaultController{})
	//微信消息回复
	beego.Router("/access", &controllers.DefaultController{})
	//文件下载
	beego.Router("/file", &controllers.FileController{})
	//社区介绍页面
	beego.Router("/introduction", &controllers.IntroductionController{})
	//狮山印象页面
	beego.Router("/impression", &controllers.ImpressionController{})
	//社区动态页面
	beego.Router("/community_service/page", &controllers.CommunityServiceController{}, "get:Page")
}
