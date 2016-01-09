package controllers

func (c *PartitionServer) API_Prof() {
	type pserver struct {
		Asd string
		Bsd string
		Csd []string
	} // 结构内的变量名首字母大写。

	u := pserver{
		"127.0.0.1:8080/A",
		"127.0.0.1:8080/B",
		[]string{"127.0.0.1:8080/qt1", "127.0.0.1:8080/qt2"},
	}
	//c.Data["json"] = &u

	c.Data["json"] = &u
	c.ServeJson()

}

/*
llers\userAPI.go:12: undefined: User

*/
