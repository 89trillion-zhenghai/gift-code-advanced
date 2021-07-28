package service

import (
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/model"
	"gift-code-Two/internal/utils"
)

func CreateAndGetGiftCode(gift model.Gift) (code string,err error) {
	//创建前先判断礼品码是否重复,重复则重新生成随机礼品码
	gCode := ""
	for {
		gCode = utils.GetGiftCode()
		if model.GiftIsExit(gCode){
			break
		}
	}
	gift.GiftCode = gCode
	//将gift以hash储存在redis里
	err = model.SetGift(gift)
	err = model.SetAvailableDetail(gift.GiftCode)
	err = model.SetAvailableTime(gift.GiftCode)
	if err != nil{
		return "",globalError.DBError("redis发送错误",globalError.RedisException)
	}
	return gCode,nil
}

func GetGiftDetail(code string) (interface{},error){
	resMap := make(map[string]interface{})
	//从MAIN_{code}获取gift主体信息
	gift, _ := model.GetGift(code)
	if len(gift) == 0 {
		return nil,globalError.GiftCodeError("礼品码有误或不存在",globalError.GiftCodeNotExist)
	}
	//从AVAILABLE_{code}获取gift领取次数
	times, _ := model.GetAvailableTime(code)
	resMap["AvailableTime"] = times
	for k, v := range gift {
		resMap[k] = v
	}
	//从DETAIL_{code}获取gift领取详情
	detail, _ := model.GetAvailableDetail(code)
	if len(detail) == 0{
		resMap["AvailableDetail"] = ""
	}else{
		resMap["AvailableDetail"] = detail
	}
	return resMap,nil
}
