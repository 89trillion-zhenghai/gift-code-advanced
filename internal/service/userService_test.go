package service

import (
	"fmt"
	"gift-code-Two/response"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
	"testing"
)

func TestLogon(t *testing.T) {
	userName := "yangzhenghai"

	got, err := Login(userName)
	if err != nil {
		log.Fatal(err.Error())
	}
	want := "新用户欢迎注册！你的通行证为："+userName
	if !reflect.DeepEqual(got,want){
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestRedeemGift(t *testing.T) {
	userName := "smallbai"
	giftCode := "33N3110J"
	gift, _ := RedeemGift(userName, giftCode)
	genera := response.GeneralReward{}
	proto.Unmarshal(gift, &genera)
	fmt.Println(gift)
}