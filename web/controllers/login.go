package controllers

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/web/middlewares"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Ctx iris.Context
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

func (login *LoginController) PostForm() mvc.Response {
	for k, v := range login.Ctx.FormValues() {
		login.Ctx.Application().Logger().Debugf("%v: %v", k, v)
	}

	token := login.Ctx.Request().Header.Get("Token")

	return mvc.Response{
		ContentType: "application/json",
		Object: iris.Map{
			"msg":   configure.USER_LOGIN_SUCCESS_STR,
			"code":  configure.CHECK_LOGIN_SUCCESS_CODE,
			"token": token,
		},
	}
}

func (login *LoginController) BeforeActivation(before mvc.BeforeActivation) {
	before.Handle("POST", "/form", "PostForm", middlewares.UserLoginCheck)
}
