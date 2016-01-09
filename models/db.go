package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/testapp/models/class"
)

func init() {
	orm.Debug = true
	switch beego.AppConfig.String("DB::db") {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			beego.AppConfig.String("DB::user"),
			beego.AppConfig.String("DB::pass"),
			beego.AppConfig.String("DB::address"),
			beego.AppConfig.String("DB::port"),
			beego.AppConfig.String("DB::dataName")))
	case "sqlite":
		orm.RegisterDriver("sqlite", orm.DR_Sqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file"))
	}

	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
}

//注册数据库驱动
//RegisterDriver(driverName,DriverType)
//RegisterDriver(驱动名称,驱动类型) 告诉 orm我要用什么数据库
//注册数据库
//RegisterDataBase(aliasName,driverName,dataSource,params...)
//RegisterDataBase(链接别名,驱动名称,数据源,可选参数链接信息...)
//注册对象模型
//RegisterModel(models) 对象模型
//开启同步
//RunSyncdb(name string,force bool,verbose bool)
//RunSyncdb(名称 )
