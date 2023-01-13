package router

import (
	ctrl "github.com/DDGRCF/DDGBlog/web/ctrl"
	helloCtrl "github.com/DDGRCF/DDGBlog/web/ctrl/api/hello"
	obbCtrl "github.com/DDGRCF/DDGBlog/web/ctrl/api/tool/obb"
	"github.com/DDGRCF/DDGBlog/web/midware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouters(app *iris.Application) {

	app.Use(midware.Cors)
	// app.Use(middlewares.GetJWT().Serve)
	// app.Handle("GET", "/hello", func(ctx iris.Context) {
	// 	ctx.JSON(
	// 		iris.Map{
	// 			"msg": conf.HELLO_STR,
	// 			"code": conf.HELLO_CODE,
	// 		})
	// })
	initAuthRouters(app)
	initIndexRouters(app)
	initLoginRouters(app)
	apiRouters(app);
}

func apiRouters(app *iris.Application) {
	apiParty := app.Party("/api")
	{
		apiParty.PartyFunc("/hello", helloCtrl.HelloHandler)

		toolsParty := apiParty.Party("/tool")
		{
			obbParty := toolsParty.Party("/obb")
			{
				obbParty.PartyFunc("/nms", obbCtrl.NmsHandler);
				obbParty.PartyFunc("/submit", obbCtrl.SubmitHandler)
			}
		}
	}
}

func initLoginRouters(app *iris.Application) {
	loginParty := app.Party("/login")
	loginFormPartyMvc := mvc.New(loginParty)
	loginFormPartyMvc.Handle(new(ctrl.LoginController))
}

func initIndexRouters(app *iris.Application) {
	indexParty := app.Party("/index")
	indexPartyMvc := mvc.New(indexParty)
	indexPartyMvc.Handle(new(ctrl.IndexController))
}

func initAuthRouters(app *iris.Application) {
	authParty := app.Party("/auth", midware.GetJWT().Serve)
	authParty.Get("", midware.DefaultJwtAuth)
}
