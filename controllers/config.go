package controllers

import (
	"blog/models"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) GetConfigPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.ConfigListPage(pages,rows)
	line := models.ConfigTotal()


	//
	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

