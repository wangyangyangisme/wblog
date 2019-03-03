package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminConfig struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ConfigKey string `json:"config_key"`
	ConfigValue string `json:"config_value"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ConfigTotal() int64 {
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminConfig))
	var res int64
	res, _ = db.Count()
	return res
}

func ConfigListPage(page int,pagesize int) []AdminConfig {
	var list []AdminConfig
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminConfig))
	conn := orm.NewCondition()

	offset := (page - 1) * pagesize
	//conn = conn.And("state", 0)

	_, _ = db.SetCond(conn).Limit(pagesize, offset).OrderBy("-created_at").All(&list)


	return list
}


func GetConfig(key ...string) interface{} {
	var res interface{}
	o := orm.NewOrm()
	db := o.QueryTable(new(AdminConfig))
	if len(key) > 1{
		var list []AdminConfig
		var keys []string
		for _,v := range key{
			keys = append(keys, v)
		}
		_,err := db.Filter("config_key__in", keys).All(&list)
		if err == orm.ErrNoRows{
			res = nil
		}
		m := make(map[string]string)
		for _,val := range list{
			m[val.ConfigKey] = val.ConfigValue
		}
		res = m
	}else{
		config := AdminConfig{}
			err := db.Filter("config_key", key[0]).One(&config)
			if err == orm.ErrNoRows{
				res = nil
			}
			res = config.ConfigValue
	}
	return res
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("prefix"),new(AdminConfig))
}
