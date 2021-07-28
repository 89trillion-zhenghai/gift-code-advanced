package ctrl

import (
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/handler"
	"gift-code-Two/internal/verify"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) (interface{},error) {
	name := c.PostForm("name")
	if !verify.ParamIsNotEmpty(name) {
		return nil, globalError.ParamError("参数为空",globalError.ParamIsEmpty)
	}
	 res, err := handler.Login(name)
	 return res,err
}

//RedeemGift 兑换礼品，返回礼品内容
func RedeemGift(c *gin.Context) (interface{},error) {
	giftCode := c.PostForm("giftCode")
	name := c.PostForm("name")
	if !verify.ParamIsNotEmpty(giftCode,name){
		return nil,globalError.ParamError("参数不能为空",globalError.ParamIsEmpty)
	}
	resMap,err := handler.RedeemGift(giftCode,name)
	return resMap,err

}

