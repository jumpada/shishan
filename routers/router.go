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
	//社区动态数据
	beego.Router("/community_service/list", &controllers.CommunityServiceController{}, "get:List")
	//社区动态详细页面
	beego.Router("/community_service/detail", &controllers.CommunityServiceController{}, "get:Detail")
	//办事指南页面
	beego.Router("/guidance/page", &controllers.GuidanceController{}, "get:Page")
	//办事指南详细页面
	beego.Router("/guidance/detail", &controllers.GuidanceController{}, "get:Detail")
	//党在身边页面
	beego.Router("/party_side/page", &controllers.PartySideController{}, "get:Page")
	//党在身边详细页面
	beego.Router("/party_side/detail", &controllers.PartySideController{}, "get:Detail")
	//党员活动页面
	beego.Router("/party_activity/page", &controllers.PartyActivityController{}, "get:Page")
	//我要咨询页面
	beego.Router("/consult/page", &controllers.ConsultController{}, "get:Page")
}
