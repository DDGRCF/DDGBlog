package configure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerCfg struct {
	Level                  string `mapstructure:"level"`
	ServerUrl              string `mapstructure:"serverUrl"`
	ServerIp               string `mapstructure:"serverIp"`
	LogDir                 string `mapstructure:"logDir"`
	LogDirSuffixTimeFormat string `mapstructure:"logDirSuffixTimeFormat"`
	SysTimeForm            string `mapstructure:"sysTimeForm"`
	SysTimeFormShort       string `mapstructure:"sysTimeFormShort"`
	CommonConfigFile       string `mapstructure:"commonConfigFile"`
	ServerConfigFile       string `mapstructure:"serverConfigFile"`
}

var defaultConf = ServerCfg{
	Level:                  "debug",
	ServerUrl:              "http://localhost",
	ServerIp:               "9999",
	LogDir:                 "log",
	LogDirSuffixTimeFormat: "2022/08",
	SysTimeForm:            "2022-08-29 15:20:05",
	SysTimeFormShort:       "2022-08-29",
	CommonConfigFile:       "configure/commonConfig.json",
	ServerConfigFile:       "configure/serverConfig.yaml",
}

var (
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
	pflag.String("level", "debug", "level to log")
	pflag.String("serverUrl", "http://localhost", "")
	pflag.String("serverIp", "9999", "")
	pflag.String("logDir", "log", "a dir to save log")
	pflag.String("logDirSuffixTimeFormat", "2022/08", "")
	pflag.String("sysTimeForm", "2022-08-29 15:20:05", "")
	pflag.String("sysTimeFormShort", "2022-08-29", "")
	pflag.String("commonConfigFile", "configure/commonConfig.json", "")
	pflag.String("serverConfigFile", "configure/serverConfig.yaml", "")
	Config.BindPFlags(pflag.CommandLine)
	Config.SetConfigFile(Config.GetString("commonConfigFile"))
	Config.AddConfigPath(".")
	err = Config.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("init pflag config failed! Code: %v\n", err))
	} else {
		Config.MergeInConfig()
	}

	replacer := strings.NewReplacer(".", "_")
	Config.SetEnvKeyReplacer(replacer)
	Config.AllowEmptyEnv(true)
	Config.AutomaticEnv()
}
