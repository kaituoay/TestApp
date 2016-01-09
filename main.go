package main

import (
	"github.com/astaxie/beego"
	_ "github.com/testapp/models"
	"github.com/testapp/models/class"
	_ "github.com/testapp/routers"
)

func main() {
	class.TestORM()
	beego.Run()

}
