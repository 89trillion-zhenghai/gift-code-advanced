package main

import (
	"gift-code-Two/app/http"
	"gift-code-Two/internal/model"
	"log"
)

func main() {
	var err error
	//MongoDB初始化
	err = model.InitMongoDB()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Redis连接初始化
	err = model.InitRc()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = http.InitRun()
	if err != nil {
		log.Println(err.Error())
	}
}



