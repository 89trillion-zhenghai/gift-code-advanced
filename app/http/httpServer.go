package http

import (
	"gift-code-Two/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRun() error {
	r := gin.Default()
	router.Route(r)
	err := r.Run(":8000")
	if err != nil{
		return err
	}
	return nil
}
