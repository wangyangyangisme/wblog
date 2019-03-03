package controllers

import (
	"blog/models"
)

type UrlController struct {
	BaseController
}

func (this *UrlController) GetUrlPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.UrlListPage(pages,rows)
	line := models.UrlTotal()


	//
	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

