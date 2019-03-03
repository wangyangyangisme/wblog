package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminMenu struct {
	Id int `json:"id"`
	Pid int `json:"pid"`
	Name string `json:"name"`
	Url string `json:"url"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminMenuShow struct {
	AdminMenu
	Child []AdminMenu
}

func MenuTotal() int64 {
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminMenu))
	var res int64
	conn := orm.NewCondition()
	//conn.And("state", 0)
	res, _ = db.SetCond(conn).Count()
	return res
}

func MenuListPage(page int,pagesize int) []AdminMenu {
	var list []AdminMenu
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminMenu))
	conn := orm.NewCondition()

	offset := (page - 1) * pagesize
	//conn = conn.And("state", 0)

	_, _ = db.SetCond(conn).Limit(pagesize, offset).OrderBy("-sort").All(&list)


	return list
}


func MenuList(page int,pagesize int) []AdminMenuShow {
	var list []AdminMenu
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminMenu))
	db.Filter("state", 0).Filter("pid", 0).All(&list)
	var res []AdminMenuShow
	for _, v := range list{
		row := AdminMenuShow{ AdminMenu:v }
		db.Filter("state", 0).Filter("pid", v.Id).OrderBy("-sort").All(&row.Child)
		res = append(res, row)
	}
	return res
}

func init(){
	orm.RegisterModelWithPrefix(beego.AppConfig.String("prefix"), new(AdminMenu))
}
