package database

import (
	"fmt"
	"time"

	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	reDB *gorm.DB
)

func GetReDB() *gorm.DB {
	return reDB
}

func CreateTable() {
	GetReDB().AutoMigrate(
		&model.User{},
	)
}

func InstanceReDB() {
	dbCfg := conf.ReDBCfg{}
	conf.CommonConfig.UnmarshalKey("DataBase.RelationDB", &dbCfg)
	dbPath := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Method, dbCfg.Host, dbCfg.Port, dbCfg.DBName)

	var err error
	reDB, err = gorm.Open(dbCfg.DBType, dbPath)

	reDB.DB().SetConnMaxLifetime(1 * time.Second)
	reDB.DB().SetConnMaxIdleTime(20)

	reDB.SingularTable(true)
	reDB.LogMode(true)
	CreateTable()
	if err != nil {
		panic(fmt.Sprintf("Create relational database fail! Code: %v", err))
	}
}
