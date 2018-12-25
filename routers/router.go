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
	//我要咨询表单
	beego.Router("/consult/form", &controllers.ConsultController{}, "get:Form")
	//提交咨询表单
	beego.Router("/consult/submit", &controllers.ConsultController{}, "get:Submit")
	//我要咨询历史记录
	beego.Router("/consult/history", &controllers.ConsultController{}, "get:History")
	//咨询历史记录列表数据
	beego.Router("/consult/list", &controllers.ConsultController{}, "get:List")
	//我要咨询详细页面
	beego.Router("/consult/detail", &controllers.ConsultController{}, "get:Detail")
	beego.Router("/voluntary/page", &controllers.VoluntaryController{}, "get:Page")
	beego.Router("/voluntary/list", &controllers.VoluntaryController{}, "get:List")
	beego.Router("/voluntary/detail", &controllers.VoluntaryController{}, "get:Detail")
	beego.Router("/vote/page", &controllers.VoteController{}, "get:Page")
	beego.Router("/vote/list", &controllers.VoteController{}, "get:List")
	beego.Router("/vote/detail", &controllers.VoteController{}, "get:Detail")
	beego.Router("/vote/submit", &controllers.VoteController{}, "get:Submit")


	beego.Router("/sign/volunteer_Info", &controllers.SignController{}, "get:VolunteerInfo")
	beego.Router("/sign/voluntary_list", &controllers.SignController{}, "get:VoluntaryList")
	beego.Router("/sign/sign", &controllers.SignController{}, "get:Sign")
}
