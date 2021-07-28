package handler

import (
	"gift-code-Two/internal/model"
	"gift-code-Two/internal/service"
)

//CreateAndGetGiftCode 创建一个礼品对象，返回一个礼品码
func CreateAndGetGiftCode(gift model.Gift) (giftCode string, err error) {
	giftCode, err = service.CreateAndGetGiftCode(gift)
	return giftCode,err
}

//GetGiftDetail 查询礼品信息
func GetGiftDetail(giftCode string)(resMap interface{},err error){
	resMap, err = service.GetGiftDetail(giftCode)
	return resMap,err
}

