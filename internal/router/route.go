package router

import (
	"gift-code-Two/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	//路由
	r.POST("/login", ctrl.Login)
	r.POST("/redeemGift", ctrl.RedeemGift)
	return r
}
