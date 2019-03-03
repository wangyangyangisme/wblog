package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int `json:"id"`
	Pid int `json:"pid"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	State int `json:"state"`
	Sort int `json:"sort"`
}

type CategoryShow struct {
	Category
	Child []Category
	Active bool
}

func CategoryListPage(page int,pagesize int) []Category {
	var list []Category
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	conn := orm.NewCondition()

	offset := (page - 1) * pagesize
	//conn = conn.And("state", 0)

	_, _ = db.SetCond(conn).Limit(pagesize, offset).OrderBy("-sort").All(&list)

	return list
}

func CategoryTotal() int64 {
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	var res int64
	conn := orm.NewCondition()
	//conn.And("state", 0)
	res, _ = db.SetCond(conn).Count()
	return res
}

func CategoryList(categoryId int) []CategoryShow {
	var list []Category
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	db.Filter("state", 0).Filter("pid", 0).OrderBy("-sort").All(&list)
	var res []CategoryShow
	for _, v := range list{
		row := CategoryShow{ Category:v }
		if categoryId == v.Id {
			row.Active = categoryId == v.Id
		}
		db.Filter("state", 0).Filter("pid", v.Id).OrderBy("-sort").All(&row.Child)
		for _,val := range row.Child{
			if categoryId == val.Id {
				row.Active = true
			}
		}
		res = append(res, row)
	}
	return res
}

func CategoryOne(id int) (Category, error) {
	var res Category
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	err := db.Filter("state", 0).Filter("id", id).One(&res)
	return res,err
}

func init () {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("prefix"), new(Category))
}