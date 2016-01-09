package controllers

import (
	"github.com/astaxie/beego"
)

type PartitionServer struct {
	beego.Controller
}

func (c *PartitionServer) Profile() {

	c.Data["电信一区"] = "127.0.0.1:8080/电信一区"
	c.Data["网通一区"] = "127.0.0.1:8080/网通一区"
	c.Data["其他服务器"] = []string{"127.0.0.1:8080/qt1区", "127.0.0.1:8080/qt2区"}

	c.TplNames = "index.html"
}
