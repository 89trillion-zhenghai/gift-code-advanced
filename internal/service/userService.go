package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gift-code-Two/internal/globalError"
	model "gift-code-Two/internal/model"
	utils "gift-code-Two/internal/utils"
	"gift-code-Two/response"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Login(name string) (user model.FindUser, err error) {
	var (
		collection = model.GetMgoCol("user", "depot")
	)
	//先判断用户是否存在
	filter := bson.M{"name": name}
	result := model.FindUser{}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		//用户不存在，新建一个用户
		user := model.InsertUser{
			Name: name,
			Depot: map[uint32]uint64{
				uint32(1001): uint64(0),
				uint32(1002): uint64(0),
			},
		}
		collection.InsertOne(context.TODO(), user)
		return model.FindUser{}, errors.New("新用户欢迎注册！你的通行证为：" + name)
	} else {
		return result, nil
	}
}

func RedeemGift(name string, code string) (resByte []byte,resErr globalError.GlobalError){
	//判断用户是否存在，不存在则新建注册一个用户
	var (
		collection = model.GetMgoCol("user", "depot")
		err        error
		uid        string
		res        *mongo.InsertOneResult
	)
	filter := bson.M{"name": name}
	result := model.FindUser{}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		//用户不存在，新建一个用户
		user := model.InsertUser{
			Name: name,
			Depot: map[uint32]uint64{
				uint32(1001): uint64(0),
				uint32(1002): uint64(0),
			},
		}
		if res, err = collection.InsertOne(context.TODO(), user); err != nil {
			log.Fatal(err)
		}
		// 取出_id
		uid = res.InsertedID.(primitive.ObjectID).Hex()
		fmt.Println(uid)
	}
	//判断礼品码是否可以兑换，需要考核题三的接口支持
	msg := utils.SendComplexGetRequest(code, name)
	if msg["message"] == nil {
		//不可兑换
		fmt.Println("不可兑换")
		return nil,globalError.GiftNotConvertible("本礼品码不可以兑换")
	}
	detail := msg["message"]["GiftDetail"]
	m := make(map[string]string)
	json.Unmarshal([]byte(detail.(string)), &m)
	if len(m) == 0 {
		//不可兑换
		fmt.Println("不可兑换")
		return nil,globalError.GiftNotConvertible("本礼品码不可以兑换")
	} else {
		//可兑换
		//获取奖品内容
		//giftDetail := msg["message"]
		//获取用户原来的物品信息
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		genera := model.GeneraReward{}
		genera.Code = int32(200)
		genera.Msg = "领取成功"
		genera.Changes = utils.MapToAnther(m)
		genera.Balance = result.Depot
		genera.Counter = utils.DepotAdd(genera.Changes, genera.Balance)
		//给用户增加奖励，mongodb更新数据
		collection.UpdateOne(context.TODO(), filter, bson.D{
			{"$set", bson.D{
				{"depot", bson.D{
					{"1001", genera.Counter[uint32(1001)]},
					{"1002", genera.Counter[uint32(1002)]},
				}},
			}},
		})
		gen := &response.GeneralReward{
			Code:    genera.Code,
			Msg:     genera.Msg,
			Changes: genera.Changes,
			Balance: genera.Balance,
			Counter: genera.Counter,
			Ext:     "扩展字段",
		}
		bytes, err := proto.Marshal(gen)
		if err != nil {
			log.Fatal(err)
		}
		return bytes,globalError.GlobalError{}
	}

}

