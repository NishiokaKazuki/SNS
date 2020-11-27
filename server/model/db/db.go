package db

import (
	"server/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	configpath = "server/config/config.toml"
)

var engine = XormConnect()

func GetDBConnect() *xorm.Engine {
	return engine
}

func XormConnect() *xorm.Engine {

	db, err := xorm.NewEngine(getDBConfig())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func getDBConfig() (string, string) {
	conf, err := config.ReadDBConfig(configpath)
	if err != nil {
		panic(err.Error())
	}

	CONNECT := conf.User + ":" + conf.Pass + "@" + conf.Protocol + "/" + conf.DbName + "?parseTime=true"

	return conf.Dbms, CONNECT
}
