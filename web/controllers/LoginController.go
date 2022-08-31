package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) Get() mvc.Result {
	// server := services.NewBookService()
	// list := service.GetList("")
	// return mvc.View{}
	return mvc.View{}
}
