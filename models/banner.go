package models

import (
	"blog/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	Id int `json:"id"`
	ArticleId int `json:"-"`
	Url string `json:"url"`
	Image string `json:"image"`
	Des string `json:"des"`
	State int `json:"state"`
}

func BannerList() []Banner {
	o := orm.NewOrm()
	db := o.QueryTable(new(Banner))
	var list []Banner
	db.Filter("state", 0).OrderBy("-id").All(&list)
	for k,v := range list{
		list[k].Image = common.Asset(v.Image)
	}
	return list
}

func BannerListPage(page int,pagesize int) []Banner {
	var list []Banner
	o := orm.NewOrm()
	db := o.QueryTable(new(Banner))
	conn := orm.NewCondition()
	offset := (page - 1) * pagesize
	conn = conn.And("state", 0)

	_, _ = db.SetCond(conn).Limit(pagesize, offset).All(&list)

	return list
}


func BannerTotal() int64 {
	o := orm.NewOrm()
	db := o.QueryTable(new(Banner))
	var res int64
	conn := orm.NewCondition()
	conn.And("state", 0)
	res, _ = db.SetCond(conn).Count()
	return res
}

func init(){
	orm.RegisterModelWithPrefix(beego.AppConfig.String("prefix"), new(Banner))
}