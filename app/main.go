package main

import (
	"gift-code-Two/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Route(r)
	r.Run(":8000")

}



