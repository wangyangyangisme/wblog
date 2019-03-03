package controllers

import (
	"blog/models"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) GetCategoryPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.CategoryListPage(pages,rows)
	line := models.CategoryTotal()

	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

