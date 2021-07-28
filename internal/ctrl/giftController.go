package ctrl

import (
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/handler"
	"gift-code-Two/internal/model"
	"gift-code-Two/internal/verify"
	"github.com/gin-gonic/gin"
	"time"
)

//CreateAndGetGiftCode 创建一个礼品对象，返回一个礼品码
func CreateAndGetGiftCode(c *gin.Context) (interface{},error){
	createName := c.PostForm("userName")
	createTime := time.Now().Format("2006-01-02 15:04:05")
	description := c.PostForm("description")
	giftType := c.PostForm("giftType")
	validity := c.PostForm("validity")
	availableTimes := c.PostForm("availableTimes")
	giftDetail := c.PostForm("giftDetail")
	if !verify.ParamIsNotEmpty(createName,description,giftType,validity,availableTimes,giftDetail){
		return nil,globalError.ParamError("参数不能为空",globalError.ParamIsEmpty)
	}
	if !verify.IsDigit(availableTimes){
		return nil,globalError.ParamError("可领取次数必须为数字",globalError.ParamIsIllegal)
	}
	if !verify.IsDigit(giftType){
		return nil,globalError.ParamError("礼品类型必须为数字",globalError.ParamIsIllegal)
	}
	tm,err := time.ParseDuration(validity)
	//过期时间 = 当前时间 + 有效期
	expireDate := time.Now().Add(tm).Unix()
	gift := model.Gift{
		CreateUser:     createName,
		CreateTime:     createTime,
		Description:    description,
		GiftType:       giftType,
		Validity:       expireDate,
		AvailableTimes: availableTimes,
		GiftDetail:     giftDetail,
	}
	giftCode, err := handler.CreateAndGetGiftCode(gift)
	return giftCode,err
}

//GetGiftDetail 查询礼品信息
func GetGiftDetail(c *gin.Context) (interface{},error){
	giftCode := c.PostForm("giftCode")
	if !verify.ParamIsNotEmpty(giftCode){
		return nil,globalError.ParamError("参数不能为空",globalError.ParamIsEmpty)
	}
	resMap, err := handler.GetGiftDetail(giftCode)
	return resMap,err
}


