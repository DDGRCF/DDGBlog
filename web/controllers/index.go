package controllers

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct {
	Ctx iris.Context
}

func (index *IndexController) Get() mvc.Result {
	appName := configure.CommonConfig.GetString("appName")

	return mvc.View{
		Name: "index/index.html",
		Data: iris.Map{
			"title": appName,
		},
	}
}
