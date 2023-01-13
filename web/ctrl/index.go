package ctrl

import (
	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct {
	Ctx iris.Context
}

func (index *IndexController) Get() mvc.Result {
	appName := conf.CommonConfig.GetString("appName")

	return mvc.View{
		Name: "index/index.html",
		Data: iris.Map{
			"title": appName,
		},
	}
}
