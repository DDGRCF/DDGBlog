package configure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DDGRCF/DDGBlog/database"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerCfg struct {
	AppName            string `mapstructure:"DDG Blog"`
	Level              string `mapstructure:"level"`
	ServerIp           string `mapstructure:"serverIp"`
	ServerPort         string `mapstructure:"serverPort"`
	LogDir             string `mapstructure:"logDir"`
	SysTimeForm        string `mapstructure:"sysTimeForm"`
	SysTimeFormShort   string `mapstructure:"sysTimeFormShort"`
	CommonConfigFile   string `mapstructure:"commonConfigFile"`
	ServerConfigFile   string `mapstructure:"serverConfigFile"`
	DatabaseConfigFile string `mapstructure:"databaseConfigFile"`
}

var defaultConf = ServerCfg{
	AppName:            "DDG Blog",
	Level:              "debug",
	ServerIp:           "127.0.0.1",
	ServerPort:         "9999",
	LogDir:             "log",
	SysTimeForm:        "2022-08-29 15:20:05",
	SysTimeFormShort:   "2022-08-29",
	CommonConfigFile:   "configure/commonConfig.json",
	ServerConfigFile:   "configure/serverConfig.yaml",
	DatabaseConfigFile: "configure/mysqlConfig.json",
}

var (
	AppLogo = `
	  _______   __       ______    _______        _______   __       ______    _______      
	/_______/\ /_/\     /_____/\ /_______/\     /_______/\ /_/\     /_____/\ /_______/\     
	\::: _  \ \\:\ \    \:::_ \ \\::: _  \ \    \::: _  \ \\:\ \    \:::_ \ \\::: _  \ \    
	 \::(_)  \/_\:\ \    \:\ \ \ \\::(_)  \/_    \::(_)  \/_\:\ \    \:\ \ \ \\::(_)  \/_   
	  \::  _  \ \\:\ \____\:\ \ \ \\::  _  \ \    \::  _  \ \\:\ \____\:\ \ \ \\::  _  \ \  
	   \::(_)  \ \\:\/___/\\:\_\ \ \\::(_)  \ \    \::(_)  \ \\:\/___/\\:\_\ \ \\::(_)  \ \ 
	    \_______\/ \_____\/ \_____\/ \_______\/     \_______\/ \_____\/ \_____\/ \_______\/
`
	Config            = viper.New()
	BypassInit string = ""
)

func init() {
	if BypassInit == "" {
		initDefault()
	}
}

func initDefault() {
	jsonBytes, err := json.Marshal(defaultConf)
	if err != nil {
		panic(fmt.Sprintf("init default config error! code: %v", err))
	}
	defaultConfig := bytes.NewReader(jsonBytes)
	viper.SetConfigType("json")
	viper.ReadConfig(defaultConfig)
	Config.MergeConfigMap(viper.AllSettings())
	pflag.String("appName", "DDG Blog", "app name")
	pflag.String("level", "debug", "level to log")
	pflag.String("serverPort", "12000", "")
	pflag.String("serverIp", "127.0.0.1", "server ip")
	pflag.String("logDir", "log", "a dir to save log")
	pflag.String("sysTimeForm", "2006-08-29 15:20:05", "")
	pflag.String("sysTimeFormShort", "2006-08-29", "")
	pflag.String("commonConfigFile", "configure/commonConfig.json", "")
	pflag.String("serverConfigFile", "configure/serverConfig.yaml", "")
	pflag.String("databaseConfigFile", "configure/mysqlConfig.json", "")
	Config.BindPFlags(pflag.CommandLine)
	Config.SetConfigFile(Config.GetString("commonConfigFile"))
	Config.AddConfigPath(".")
	err = Config.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("init pflag config failed! Code: %v\n", err))
	} else {
		Config.MergeInConfig()
	}
	database.MysqlInit(Config.GetString("databaseConfigFile"))
	replacer := strings.NewReplacer(".", "_")
	Config.SetEnvKeyReplacer(replacer)
	Config.AllowEmptyEnv(true)
	Config.AutomaticEnv()
}
