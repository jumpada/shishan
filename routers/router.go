package routers

import (
	"github.com/astaxie/beego"
	"shishan/controllers"
)

func init() {
	//微信接入验证
	beego.Router("/access", &controllers.DefaultController{}, "get:Access")
	//社区概况页面
	beego.Router("/overview/page", &controllers.OverviewController{}, "get:Page")
	//社区资讯页面
	beego.Router("/info/page", &controllers.InfoController{}, "get:Page")
	beego.Router("/info/list", &controllers.InfoController{}, "get:List")
}
