package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/DDGRCF/myBlob/configure"
	"github.com/DDGRCF/myBlob/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kr/pretty"
)

func main() {
	// 创建Iris框架
	app := iris.New()
	app.Logger().SetLevel(configure.Config.GetString("level"))
	app.Logger().SetOutput(io.MultiWriter(func() *os.File {
		dirName := path.Join(configure.Config.GetString("logDir"),
			time.Now().Format(configure.Config.GetString("logDirSuffixTimeFormat")))
		os.MkdirAll(dirName, os.ModePerm)
		today := time.Now().Format(configure.Config.GetString("sysTimeForm"))
		fileName := path.Join(dirName, today+".log")
		fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("open file %v failed! Code: %v", fileName, err))
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
	// app.RegisterView(iris.HTML("./web/views", ".html"))
	app.HandleDir("/content", "./web/content")
	var config iris.Configurator
	serverConfigFile := configure.Config.GetString("serverConfigFile")
	if exist, err := utils.PathIsExists(serverConfigFile); !exist {
		config = iris.WithConfiguration(iris.Configuration{
			FireMethodNotAllowed: true,
			DisableStartupLog:    true,
			EnableOptimizations:  true,
			Charset:              "UTF-8",
		})
		app.Logger().Warnf("load server config failed! Code: %v! use default server config", err)
	} else {
		config = iris.WithConfiguration(iris.YAML(serverConfigFile))
	}

	cfg := configure.ServerCfg{}
	configure.Config.Unmarshal(&cfg)
	app.Logger().Infof(`
				_    
		 / ___| ___  
		\ |  _ / _ \ 
		/ |_| | (_) |
		\____|\___/	
	`)
	app.Logger().Debugf("load the server config: %#v", pretty.Formatter(cfg))
	app.Run(
		iris.Addr(":"+configure.Config.GetString("ServerIp")),
		iris.WithoutServerError(iris.ErrServerClosed),
		config,
	)
}
