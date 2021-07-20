package main

import (
	"fmt"
	"gift-code-Two/app/http"
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/model"
	"log"
)

func main() {
	var err error

	err = model.InitMongoDB()
	if err != nil {
		log.Println(globalError.MongoDBError("数据库连接失败！请检查并重试"))
		return
	}
	err = http.InitRun()
	if err != nil {
		fmt.Println(err.Error())
	}
}



