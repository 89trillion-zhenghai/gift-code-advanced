package ctrl

import (
	"gift-code-Two/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.Query("name")
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
	name := c.Query("name")
	code := c.Query("giftCode")
	bytes := service.RedeemGift(name, code)
	c.JSON(http.StatusOK, gin.H{
		"message": bytes,
	})
}

