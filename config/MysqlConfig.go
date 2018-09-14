package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

func Init() (*xorm.Engine) {
	var err error
	DBEngine, err := xorm.NewEngine("mysql", "autoDev:auto2015@tcp(10.0.3.200)/virtualNumber?charset=utf8")
	if err != nil {
		panic(err)
	}
	DBEngine.SetMaxIdleConns(30)
	DBEngine.SetMaxOpenConns(50)
	DBEngine.ShowSQL(true)
	return DBEngine
}
