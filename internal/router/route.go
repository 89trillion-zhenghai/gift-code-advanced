package router

import (
	"gift-code-Two/internal/ctrl"
	"gift-code-Two/internal/globalError"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	//路由转发
	r.POST("/createAndGetGiftCode",globalError.ErrorHandler(ctrl.CreateAndGetGiftCode))
	r.POST("/getGiftDetail",globalError.ErrorHandler(ctrl.GetGiftDetail))
	r.POST("/login", globalError.ErrorHandler(ctrl.Login))
	r.POST("/redeemGift", globalError.ErrorHandler(ctrl.RedeemGift))
	return r
}
