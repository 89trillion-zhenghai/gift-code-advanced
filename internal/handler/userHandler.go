package handler

import "gift-code-Two/internal/service"

func Login(name string) (interface{},error) {
	user, err := service.Login(name)
	return user,err
}

//RedeemGift 兑换礼品，返回礼品内容
func RedeemGift(giftCode string,name string) (resMap interface{},err error){
	resMap,err = service.RedeemGift(giftCode,name)
	return resMap,err
}
