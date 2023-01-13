package ctrl

import (
	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/web/midware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Ctx iris.Context
}

func (login *LoginController) Get() mvc.Result {
	appName := conf.ServerConfig.GetString("appName")
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
			"msg":   conf.USER_LOGIN_SUCCESS_STR,
			"code":  conf.CHECK_LOGIN_SUCCESS_CODE,
			"token": token,
		},
	}
}

func (login *LoginController) BeforeActivation(before mvc.BeforeActivation) {
	before.Handle("POST", "/form", "PostForm", midware.UserLoginCheck)
}
