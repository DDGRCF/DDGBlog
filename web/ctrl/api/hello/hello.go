package ctrl

import (
	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/model"
	"github.com/kataras/iris/v12"
)


func HelloHandler(hello iris.Party) {
	hello.Get("/", func (ctx iris.Context) {
		response := model.Response{
			Msg: conf.SUCCESS_CODE_STR,
			Code: conf.SUCCESS_CODE,
			Data: "Hello World",
		}
		ctx.JSON(response)
	})
}