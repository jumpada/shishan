package routers

import (
	"github.com/astaxie/beego"
	"shishan/controllers"
)

func init() {
	//微信接入验证
	beego.Router("/access", &controllers.DefaultController{}, "get:Access")
	//社区介绍页面
	beego.Router("/intro", &controllers.IntroController{})
	//社区资讯页面
	beego.Router("/info/page", &controllers.InfoController{}, "get:Page")
	beego.Router("/info/list", &controllers.InfoController{}, "get:List")
}
