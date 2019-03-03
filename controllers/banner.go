package controllers

import (
	"blog/models"
)

type BannerController struct {
	BaseController
}

func (this *BannerController) GetBannerPage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.BannerListPage(pages,rows)
	line := models.BannerTotal()

	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,lists}

	this.Data["json"] = &j
	this.ServeJSON()
 }

