package ctrl

import (
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.PostForm("name")
	if len(name) == 0 {
		err := globalError.Param("参数为空")
		c.JSON(err.Status,err)
		return
	}

	if res, err := service.Login(name); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
}

func RedeemGift(c *gin.Context) {
	name := c.PostForm("name")
	code := c.PostForm("giftCode")
	if len(name) == 0 {
		err := globalError.Param("参数为空")
		c.JSON(err.Status,err)
		return
	}
	bytes,err := service.RedeemGift(name, code)
	if len(err.Error())!=0 {
		c.JSON(err.Status, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": bytes,
		})
	}

}

