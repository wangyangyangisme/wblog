package controllers

import (
	"blog/models"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) GetArticlePage()  {
	rows,_ := this.GetInt("rows")
	pages,_ := this.GetInt("page")


	lists := models.ArticleListByTime(pages,rows,"")
	line := models.ArticleTotal(0,"")

	var articles []interface{}

	for _,v := range lists {
		articles = append(articles,v.Article)
	}

	//
	type Json struct {
		Total int64 `json:"total"`
		Rows interface{} `json:"rows"`
	}

	j := Json{line,articles}

	this.Data["json"] = &j
	this.ServeJSON()
 }

