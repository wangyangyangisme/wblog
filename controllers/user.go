package controllers

import (
	"blog/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) GetUserPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.AdminUserListPage(pages,rows)
	line := models.AdminUserTotal()


	//
	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

