package database

import (
	"fmt"

	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var (
	unReDB *redis.Database
)

func GetUnReDB() *redis.Database {
	return unReDB
}

func InstanceUnReDB() {
	dbCfg := configure.UnReDBCfg{}
	configure.CommonConfig.UnmarshalKey("DataBase.UnRelationDB", &dbCfg)
	unReDB = redis.New(redis.Config{
		Network:   dbCfg.Method,
		Addr:      fmt.Sprintf("%v:%v", dbCfg.Host, dbCfg.Port),
		Password:  dbCfg.Password,
		Database:  "",
		MaxActive: 10,
		Timeout:   redis.DefaultRedisTimeout,
		Prefix:    dbCfg.Prefix,
	})
}
