package routers

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/web/controllers"
	"github.com/DDGRCF/DDGBlog/web/middlewares"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouters(app *iris.Application) {

	app.Use(middlewares.Cors)
	// app.Use(middlewares.GetJWT().Serve)
	app.Handle("GET", "/hello", func(ctx iris.Context) {
		ctx.JSON(
			iris.Map{
				"msg": configure.HELLO_STR,
				"code": configure.HELLO_CODE,
			})
	})
	initAuthRouters(app)
	initIndexRouters(app)
	initLoginRouters(app)
}

func initLoginRouters(app *iris.Application) {
	loginParty := app.Party("/login")
	loginFormPartyMvc := mvc.New(loginParty)
	loginFormPartyMvc.Handle(new(controllers.LoginController))
}

func initIndexRouters(app *iris.Application) {
	indexParty := app.Party("/index")
	indexPartyMvc := mvc.New(indexParty)
	indexPartyMvc.Handle(new(controllers.IndexController))
}

func initAuthRouters(app *iris.Application) {
	authParty := app.Party("/auth", middlewares.GetJWT().Serve)
	authParty.Get("", middlewares.DefaultJwtAuth)
}
