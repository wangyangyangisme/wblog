package controllers

import (
	"blog/models"
)

type MenuController struct {
	BaseController
}

func (this *MenuController) GetMenuPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.MenuListPage(pages,rows)
	line := models.MenuTotal()


	//
	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

