package handler

import (
	"fmt"
	"gift-code-Two/internal/model"
	"gift-code-Two/internal/service"
	"gift-code-Two/response"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
)


func TestCreateAndGetGiftCode(t *testing.T) {
	model.InitRc()
	gift := model.Gift{
		CreateUser:  "admin",
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Description: "十周年活动奖励",
		GiftType:    "2",
		AvailableTimes: "20",
		GiftDetail:  "{\"1001\":\"10\",\"1002\":\"5\"}",
	}
	validity := "10m"
	tm,err := time.ParseDuration(validity)
	//过期时间 = 当前时间 + 有效期
	expireDate := time.Now().Add(tm).Unix()
	gift.Validity = expireDate
	code, err := service.CreateAndGetGiftCode(gift)
	if err != nil {
		t.Error(err)
	}
	println(code)
}

func TestGetGiftDetail(t *testing.T) {
	model.InitRc()
	giftCode := "63HF5Y63"
	detail, err := service.GetGiftDetail(giftCode)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(detail)
}

func TestLogin(t *testing.T) {
	model.InitMongoDB()
	name := "yangzhenghai"
	login, err := service.Login(name)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(login)
}

func TestRedeemGift(t *testing.T) {
	model.InitRc()
	model.InitMongoDB()
	code := "63HF5Y63"
	name := "yangzhenghai"
	bytes, err := service.RedeemGift(code, name)
	if err != nil{
		t.Error(err)
	}
	gen := response.GeneralReward{}
	err = proto.Unmarshal(bytes, &gen)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(gen)
}
