package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id int `json:"id"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
	Account string `json:"account"`
	Password string `json:"password"`
	ClearPassword string `json:"clear_password"`

}


func AdminUserTotal() int64 {
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminUser))
	var res int64
	res, _ = db.Count()
	return res
}


func AdminUserListPage(page,pagesize int) []AdminUser{
	var list []AdminUser
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminUser))
	offset := (page - 1) * pagesize
	db.Limit(pagesize, offset).OrderBy("-account").All(&list)

	return list
}

func init(){
	orm.RegisterModelWithPrefix(beego.AppConfig.String("prefix"), new(AdminUser))
}