package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	beego.Router("/category/:id:int", &controllers.IndexController{},"get:List")
	beego.Router("/article/:id:int", &controllers.IndexController{},"get:Article")
	beego.Router("/tag", &controllers.IndexController{},"get:Tag")
	beego.Router("/news", &controllers.IndexController{},"get:News")
	beego.Router("/crawler", &controllers.CrawlerController{},"get:Index")



	beego.Router("/admin", &controllers.IndexController{},"get:Admin")
	beego.Router("/settings", &controllers.IndexController{},"get:Settings")
	beego.Router("/home", &controllers.IndexController{},"get:Home")
	beego.Router("/articlemanage", &controllers.IndexController{},"get:ArticleManage")
	beego.Router("/bannermanage", &controllers.IndexController{},"get:BannerManage")
	beego.Router("/categorymanage", &controllers.IndexController{},"get:CategoryManage")
	beego.Router("/menumanage", &controllers.IndexController{},"get:MenuManage")
	beego.Router("/otherarticlemanage", &controllers.IndexController{},"get:OtherArticleManage")
	beego.Router("/configmanage", &controllers.IndexController{},"get:ConfigManage")
	beego.Router("/usermanage", &controllers.IndexController{},"get:UserManage")
	beego.Router("/articleform", &controllers.IndexController{},"get:ArticleForm")
	beego.Router("/login", &controllers.IndexController{},"get:Login")



	beego.Router("/articlelist", &controllers.ArticleController{},"get:GetArticlePage")

	beego.Router("/bannerlist", &controllers.BannerController{},"get:GetBannerPage")

	beego.Router("/categorylist", &controllers.CategoryController{},"get:GetCategoryPage")

	beego.Router("/menulist", &controllers.MenuController{},"get:GetMenuPage")

	beego.Router("/urllist", &controllers.UrlController{},"get:GetUrlPage")

	beego.Router("/configlist", &controllers.ConfigController{},"get:GetConfigPage")

	beego.Router("/userlist", &controllers.UserController{},"get:GetUserPage")
}
