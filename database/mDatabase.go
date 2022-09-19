package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var Db *sqlx.DB

type MysqlConfig struct {
	MysqlIp   string `mapstructure:"mysqlIp"`
	MysqlPort string `mapstructure:"mysqlPort"`
}

func MysqlInit(configFile string) {
	vip := viper.New()
	vip.AddConfigPath(".")
	vip.SetConfigFile(configFile)
	if err := vip.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("parse the mysql config failed! Code: %v", err))
	}
	mysqlConfig := MysqlConfig{}
	if err := vip.Unmarshal(&mysqlConfig); err != nil {
		panic(fmt.Sprintf("unmarshal the config failed! Code: %v", err))
	}

	mysqlConn := "root:9696@tcp(" + mysqlConfig.MysqlIp + ":" + mysqlConfig.MysqlPort + ")/test"
	fmt.Println(mysqlConn)
	database, err := sqlx.Open("mysql", mysqlConn)
	if err != nil {
		panic(fmt.Sprintf("create the database failed! Code: %v", err))
	}
	Db = database
	defer database.Close()
}
