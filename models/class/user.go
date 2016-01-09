package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:"pk"`
	Nick  string
	Email string
	Test  string
}

func TestORM() {

	o := orm.NewOrm()
	u := User{"jike", "geek", "adsd@qq.c", "asdasd"}
	o.Insert(&u)
	u1 := User{Id: "jike"}
	o.Read(&u1)
	fmt.Println(u1)

	u1.Nick = "lisan"
	o.Update(&u)

	u2 := User{Id: "jike"}
	o.Read(&u2)
	fmt.Println(u2)
	o.Delete(&u2)

	//	o.Read(&u2)
	//	fmt.Println(u2)
}
