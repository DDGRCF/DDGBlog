package routers

import (
	"github.com/DDGRCF/DDGBlog/web/controllers"
	"github.com/DDGRCF/DDGBlog/web/middlewares"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func InitRouters(app *iris.Application) {

	sessionID := "mSession"
	session := sessions.New(sessions.Config{
		Cookie: sessionID,
	})

	app.Use(middlewares.Cors)
	app.Use(middlewares.GetJWT().Serve)
	initIndexRouters(app, session)
	initLoginRouters(app, session)
}

func initLoginRouters(app *iris.Application, session *sessions.Sessions) {
	loginParty := app.Party("/login")
	loginFormPartyMvc := mvc.New(loginParty)
	loginFormPartyMvc.Register(session.Start)
	loginFormPartyMvc.Handle(new(controllers.LoginController))
}

func initIndexRouters(app *iris.Application, session *sessions.Sessions) {
	indexParty := app.Party("/index")
	indexPartyMvc := mvc.New(indexParty)
	indexPartyMvc.Register(session.Start)
	indexPartyMvc.Handle(new(controllers.IndexController))

}
