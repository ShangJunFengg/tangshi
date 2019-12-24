package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"reflect"
)

type DataBases struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &DataBases{}

type Server struct {
	HttpPort string
}

var ServerSetting = &Server{}

func InitConf() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("load app.ini error : %v", err)
	}
	err1 := cfg.Section("database").MapTo(DatabaseSetting)
	err2 := cfg.Section("server").MapTo(ServerSetting)
	SmartPrint(*DatabaseSetting)
	SmartPrint(*ServerSetting)
	if err1 == nil && err2 == nil {
		log.Println("--success ini--")
	}
}

func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Print("  ")
	}
	fmt.Println()
}
