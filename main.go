package main

import (
	"MyTang/db"
	"MyTang/routers"
	"MyTang/setting"
	"log"
)

func main() {
	log.Printf("--start server--")
	log.Printf("--start ini--")
	setting.InitConf()
	log.Printf("--start db--")
	//db.InitDb()
	db.InitXorm()
	log.Printf("--start router--")
	routers.InitRouter()

}
