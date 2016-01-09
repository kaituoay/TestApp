package routers

import (
	"github.com/astaxie/beego"
	"github.com/testapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/PartitionServer", &controllers.PartitionServer{}, "get:Profile")
	beego.Router("/api/PartitionServer", &controllers.PartitionServer{}, "get:API_Prof")
	//  beego.Router("rootpath",&controllers , ...)
	//将 域名的根目录映射给了 MainController方法
}
