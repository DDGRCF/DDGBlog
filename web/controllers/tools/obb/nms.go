package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ToolsObbNmsController struct {
	Ctx iris.Context
}

func (obbNms *ToolsObbNmsController) GetToolsObbNms() mvc.Response {
	return mvc.Response{}
}