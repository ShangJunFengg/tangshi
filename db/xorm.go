package db

import (
	"MyTang/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var Engine *xorm.Engine

func InitXorm() {

	var err error
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Print("init Xorm error ")
	}
	errPing := Engine.Ping()
	if errPing == nil {
		log.Println("success DB")
	}
	Engine.ShowSQL(true)
}
