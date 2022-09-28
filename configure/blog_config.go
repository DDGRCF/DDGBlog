package configure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CommonCfg struct {
	AppName          string `json:"appName"`
	ServerIP         string `json:"serverIP"`
	ServerPort       int    `json:"serverPort"`
	SysTimeForm      string `json:"sysTimeForm"`
	SysTimeFormShort string `json:"sysTimeFormShort"`
	CommonConfigFile string `json:"commonConfigFile"`
}

type ServerCfg struct {
	FireMethodNotAllowed bool   `json:"fireMethodNotAllow"`
	DisableStartupLog    bool   `json:"disableStartupLog"`
	EnableOptimizations  bool   `json:"enableOptimizations"`
	CharSet              string `json:"charSet"`
}

type ReDBCfg struct {
	Method   string `json:"method"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	DBType   string `json:"dbType"`
}

type UnReDBCfg struct {
	Method   string `json:"method"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
	DBType   string `json:"dbType"`
}

type LogCfg struct {
	Dir   string `json:"dir"`
	Level string `json:"level"`
}

var defServerConfig = ServerCfg{
	FireMethodNotAllowed: true,
	DisableStartupLog:    true,
	EnableOptimizations:  true,
	CharSet:              "UTF-8",
}

var defLogConfig = LogCfg{
	Dir:   "log",
	Level: "debug",
}

var defCommonConfig = CommonCfg{
	AppName:          "DDG Blog",
	ServerIP:         "127.0.0.1",
	ServerPort:       9999,
	SysTimeForm:      "2022-08-29 15:20:05",
	SysTimeFormShort: "2022-08-29",
	CommonConfigFile: "configure/common_config.json",
}

var defReDBConfig = ReDBCfg{
	Method:   "tcp",
	Host:     "127.0.0.1",
	Port:     3306,
	User:     "root",
	Password: "",
	DBName:   "DDGBlogDB",
	DBType:   "mysql",
}

var defUnReDBConfig = UnReDBCfg{
	Method:   "tcp",
	Host:     "127.0.0.1",
	Port:     6379,
	Password: "",
	Prefix:   "",
	DBType:   "redis",
}

var (
	AppLogo = `
		 ______   ______   _______  ______   _        _______  _______ 
		(  __  \ (  __  \ (  ____ \(  ___ \ ( \      (  ___  )(  ____ \
		| (  \  )| (  \  )| (    \/| (   ) )| (      | (   ) || (    \/
		| |   ) || |   ) || |      | (__/ / | |      | |   | || |      
		| |   | || |   | || | ____ |  __ (  | |      | |   | || | ____ 
		| |   ) || |   ) || | \_  )| (  \ \ | |      | |   | || | \_  )
		| (__/  )| (__/  )| (___) || )___) )| (____/\| (___) || (___) |
		(______/ (______/ (_______)|/ \___/ (_______/(_______)(_______)
`
	CommonConfig = viper.New()
	ServerConfig = viper.New()
	ReDBConfig   = viper.New()
	UnReDBConfig = viper.New()
	LogConfig    = viper.New()
)

func init() {
	initDefault()
}

// Load config prior : default -> command -> config -> env
func initDefault() {
	commonJsonBytes, err := json.Marshal(defCommonConfig)
	if err != nil {
		panic(fmt.Sprintf("init server default config error! code: %v", err))
	}
	_commonConfig := bytes.NewReader(commonJsonBytes)
	viper.SetConfigType("json")
	viper.ReadConfig(_commonConfig)
	CommonConfig.MergeConfigMap(viper.AllSettings())
	pflag.String("appName", "DDG Blog", "app name")
	pflag.Int("serverPort", 12000, "")
	pflag.String("serverIp", "127.0.0.1", "server ip")
	pflag.String("sysTimeForm", "2006-08-29 15:20:05", "")
	pflag.String("sysTimeFormShort", "2006-08-29", "")
	pflag.String("commonConfigFile", "configure/common_config.json", "")
	CommonConfig.BindPFlags(pflag.CommandLine)
	CommonConfig.SetConfigFile(CommonConfig.GetString("commonConfigFile"))
	CommonConfig.AddConfigPath(".")
	err = CommonConfig.ReadInConfig()
	if err != nil {
		log.Println("Read common config fail! Will load server default config")
	} else {
		CommonConfig.MergeInConfig()
	}

	replacer := strings.NewReplacer(" ", "_") // replace
	CommonConfig.SetEnvKeyReplacer(replacer)
	CommonConfig.BindEnv("mode")
	CommonConfig.SetEnvPrefix(ServerConfig.GetString("appName"))

	CommonConfig.AllowEmptyEnv(true)
	CommonConfig.AutomaticEnv()

	if !CommonConfig.IsSet("log") {
		log.Println("Find log config fail! Use default config")
		CommonConfig.SetDefault("Log", defLogConfig)
	}
	if !CommonConfig.IsSet("server") {
		log.Println("Find server config fail! Use default config")
		CommonConfig.SetDefault("Server", defServerConfig)
	}
	if !CommonConfig.IsSet("database") {
		log.Println("Find database config fail! Use default config")
		CommonConfig.SetDefault("database.relationDB", defReDBConfig)
		CommonConfig.SetDefault("database.unrelationDB", defUnReDBConfig)
	} else {
		if !CommonConfig.IsSet("database.relationDB") {
			log.Println("Find relational database config fail! Use default config")
			CommonConfig.SetDefault("database.relationDB", defReDBConfig)
		}
		if !CommonConfig.IsSet("database.unrelationDB") {
			log.Println("Find unrelational database config fail! Use default config")
			CommonConfig.SetDefault("database.unrelationDB", defUnReDBConfig)
		}
	}

	dump.P(CommonConfig.AllSettings())
}
