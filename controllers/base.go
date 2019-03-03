package controllers

import "github.com/astaxie/beego"
import "github.com/astaxie/beego/context"


type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare()  {
	uinfo := this.Ctx.Input.Session("userinfo")
	if uinfo == nil {
		this.Ctx.Redirect(301, "/login")
		//return
		println("redirect")
	} else {
		println("========")
		//return
	}
}

func init()  {
	//CheckAuth()
}

func CheckAuth() {

	var check = func(ctx *context.Context) {
		uinfo := ctx.Input.Session("userinfo")
		if uinfo == nil {
			ctx.Redirect(302, "/login")
			return
		} else {
			return
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, check)
}