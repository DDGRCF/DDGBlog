package controllers

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) Get() mvc.Result {
	appName := configure.Config.GetString("appName")

	return mvc.View{
		Name: "login/login.html",
		Data: iris.Map{
			"title": appName,
		},
	}
}

func (c *LoginController) Post(ctx iris.Context) mvc.Result {
	emailAddr := ctx.FormValue("floatingInput")
	password := ctx.FormValue("floatingPassword")

	ctx.Application().Logger().Infof("EmailAddr: %v, Password: %v\n", emailAddr, password)

	return mvc.Response{
		ContentType: "text/html",
		Code:        200,
		Text:        "<h1>Weclome</h1>",
	}
}
