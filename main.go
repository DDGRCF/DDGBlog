package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/database"
	"github.com/DDGRCF/DDGBlog/routers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// 初始数据库
	database.InstanceReDB()
	database.InstanceUnReDB()

	// 初始化app
	app := iris.New()

	logCfg := configure.LogCfg{}
	configure.CommonConfig.UnmarshalKey("Log", &logCfg)

	// 初始化server level
	app.Logger().SetLevel(logCfg.Level)
	app.Logger().SetOutput(io.MultiWriter(func() *os.File {
		dirName := path.Join(logCfg.Dir,
			time.Now().Format(configure.ServerConfig.GetString("sysTimeFormShort")))
		os.MkdirAll(dirName, os.ModePerm)
		today := time.Now().Format(configure.ServerConfig.GetString("sysTimeForm"))
		fileName := path.Join(dirName, today+".log")
		fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Open file %v failed! Code: %v", fileName, err))
		}
		return fd
	}(), os.Stdout))

	requestLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	})

	app.Use(requestLogger)
	app.Use(recover.New())
	// app.HandleDir("/content", "./web/content")
	// pugEngine := iris.Django("./web/views", ".html")
	// DDGBLOG_MODE="development" 要全部大写
	// if configure.CommonConfig.GetString("mode") == "development" {
	// 	app.Logger().Debug("Load dynamic template views")
	// 	pugEngine.Reload(true)
	// }

	// app.RegisterView(pugEngine)
	var irisConfig iris.Configurator
	_irisConfig := iris.Configuration{}
	configure.CommonConfig.UnmarshalKey("server", &_irisConfig)
	irisConfig = iris.WithConfiguration(_irisConfig)

	routers.InitRouters(app)
	app.Logger().Info(configure.AppLogo)
	app.Run(
		iris.Addr(fmt.Sprintf("%s:%s",
			configure.CommonConfig.GetString("serverIp"),
			configure.CommonConfig.GetString("serverPort"))),
		iris.WithoutServerError(iris.ErrServerClosed),
		irisConfig,
	)
}
