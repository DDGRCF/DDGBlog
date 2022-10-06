package controllers

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/web/middlewares"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type LoginController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (login *LoginController) Get() mvc.Result {
	appName := configure.ServerConfig.GetString("appName")
	return mvc.View{
		Name: "login/login.html",
		Data: iris.Map{
			"title": appName,
		},
	}
}

func (login *LoginController) PostForm() mvc.Result {
	for k, v := range login.Ctx.FormValues() {
		login.Ctx.Application().Logger().Debugf("%v: %v", k, v)
	}
	return mvc.Response{
		ContentType: "text/html",
		Code:        200,
		Text:        "<h1 align=\"center\" >Weclome</h1>",
	}
}

func (login *LoginController) BeforeActivation(before mvc.BeforeActivation) {
	before.Handle("POST", "/form", "PostForm", middlewares.BeforePostFormCheckFormat)
}
