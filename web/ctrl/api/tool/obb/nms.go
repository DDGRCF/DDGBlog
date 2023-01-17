package ctrl

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/model"
	"github.com/kataras/iris/v12"
)


func NmsHandler(nms iris.Party) {
	nms.Post("/upload", 
		iris.LimitRequestBodySize(conf.NMS_UPLOAD_MAX_SIZE << 20), // 1 << 20 = 1MB
		uploadHandler)
	nms.Get("/maxsize", getMaxSize)
}

func getMaxSize(ctx iris.Context) {
	response := model.Response{
		Msg: conf.SUCCESS_CODE_STR,
		Code: conf.SUCCESS_CODE,
		Data: iris.Map {
			"size": conf.NMS_UPLOAD_MAX_SIZE,
			"unit": "MB",
		},
	}

	ctx.JSON(response);
}

func uploadHandler(ctx iris.Context) {
	_, ret, err := ctx.UploadFormFiles(conf.GetUploadPath(), beforeFileSave)
	ctx.Application().Logger().Println(ctx.FormValue("scoreThr")) // 这个能够获得值
	ctx.Application().Logger().Println(ctx.FormValue("iouThr"))
	var response model.Response
	if (err != nil) {
		ctx.Application().Logger().Errorf("Recv %v files, %v", ret, err)
		response = model.Response{
			Msg: conf.FAILURE_CODE_STR,
			Code: conf.FAILURE_CODE,
			Data: fmt.Sprintf("上传失败, %v", err),
		}
	} else {
		ctx.Application().Logger().Infof("Recv %v files", ret);
		response = model.Response{
			Msg: conf.SUCCESS_CODE_STR,
			Code: conf.SUCCESS_CODE,
		}
	}
	ctx.JSON(response)
}

func beforeFileSave(ctx iris.Context, file *multipart.FileHeader) bool {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1);
	ip = strings.Replace(ip, ":", "_", -1);
	file.Filename = ip + "-" + file.Filename
	return true
}