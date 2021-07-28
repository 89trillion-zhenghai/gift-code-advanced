package service

import (
	"encoding/json"
	"gift-code-Two/internal/globalError"
	"gift-code-Two/internal/model"
	"gift-code-Two/internal/utils"
	"gift-code-Two/response"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
	"time"
)

func Login(name string) (interface{}, error) {
	//先判断用户是否存在
	res, exit := model.UserIsExit(name)
	if exit {
		return res,nil
	}
	//用户不存在，注册一个新用户
	user := model.InsertUser{
		Name: name,
		Depot: map[uint32]uint64{
			uint32(1001): uint64(0),
			uint32(1002): uint64(0),
		},
	}
	err := model.Register(user)
	if err != nil {
		return nil,err
	}
	return nil,globalError.Register(name)
}

func RedeemGift(code string, name string) ([]byte,error){
	//判断该用户是否存在
	user, exit := model.UserIsExit(name)
	if !exit {
		return nil,globalError.UserError("用户不存在")
	}
	//判断礼品码是否存在
	gift, _ := model.GetGift(code)
	if len(gift) == 0 {
		return nil,globalError.GiftCodeError("礼品码有误或不存在",globalError.GiftCodeNotExist)
	}
	//判断礼品码是否失效
	now := time.Now().Unix()
	validity := gift["Validity"]
	val, _ := strconv.Atoi(validity)
	if int64(val) < now{
		return nil,globalError.GiftCodeError("礼品码失效",globalError.GiftCodeExpired)
	}
	//获取礼品码的种类
	giftType := gift["GiftType"]
	switch giftType {
	case "1":
		//第一类礼品码：指定用户一次性消耗
		//只需要判断领取列表是否有值，没有值就直接领取
		flag := model.GiftIsAvailed(code)
		if !flag {
			return nil,globalError.GiftCodeError("礼品码已失效",globalError.GiftCodeIsInvalid)
		}
		//领取操作
		res, err := receive(code, name, user, gift)
		if err != nil {
			return nil,err
		}
		return res,nil
	case "2":
		//第二类礼品码：不指定用户限制兑换次数
		//判断可领取次数是否大于0
		availableTime, _ := model.GetAvailableTime(code)
		time, _ := strconv.Atoi(gift["AvailableTimes"])
		times,_ := strconv.Atoi(availableTime)
		if time - times < 1 {
			return nil,globalError.GiftCodeError("礼品已被领取光了",globalError.GiftIsOver)
		}
		//判断是否领取过
		if !model.UserIsAvailed(code,name){
			return nil,globalError.GiftCodeError("你已经领取过本礼品了",globalError.GiftCodeReceived)
		}
		//领取操作
		res, err := receive(code, name, user, gift)
		if err != nil {
			return nil,err
		}
		return res,nil
	default:
		//第三类礼品码：不限用户不限次数兑换
		//判断是否领取过
		if !model.UserIsAvailed(code,name){
			return nil,globalError.GiftCodeError("你已经领取过本礼品了",globalError.GiftCodeReceived)
		}
		//领取操作
		res, err := receive(code, name, user, gift)
		if err != nil {
			return nil,err
		}
		return res,nil
	}
}

//开启事务操作  redis中的写操作 1、领取次数+1 2、领取列表添加，mongodb中的写操作 1、给用户增加奖励，三步操作必须保证全部成功或者有一个失败就回退所有操作
func receive(code string ,name string,user model.FindUser,gift map[string]string) ([]byte,error) {
	//领取次数+1，领取列表追加
	err := model.IncrAvailableAndAppendUser(code, name)
	if err != nil {
		return nil,err
	}
	//mongodb给用户增加奖励,用户原来的信息：user 礼包内容：gift["GiftDetail"]
	change := make(map[string]string)
	json.Unmarshal([]byte(gift["GiftDetail"]), &change)
	changes := utils.MapToAnther(change)
	balance := make(map[uint32]uint64)
	for key := range changes {
		balance[key] = user.Depot[key]
	}
	result := utils.DepotAdd(user.Depot,changes)
	//给用户增加奖励，mongodb更新数据
	err = model.UpdateUser(result, name)
	if err!= nil {
		//回滚mongodb和redis的操作
		model.Rollback(code,user)
		return nil,globalError.ServerError("服务器异常,请重试！")
	}
	genera := &response.GeneralReward{
		Code:    int32(200),
		Msg:     "领取成功",
		Changes: changes,
		Balance: balance,
		Counter: utils.DepotAdd(changes, balance),
		Ext:     "扩展字段",
	}
	bytes, err := proto.Marshal(genera)
	if err != nil {
		log.Fatal(err)
	}
	return bytes,nil
}


