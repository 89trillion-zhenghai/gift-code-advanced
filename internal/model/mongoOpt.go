package model

import (
	"context"
	"gift-code-Two/internal/globalError"
	"go.mongodb.org/mongo-driver/bson"
)

//UserIsExit 判断用户是否存在 存在返回用户信息，不存在返回false
func UserIsExit(name string) (FindUser,bool) {
	filter := bson.M{"name": name}
	result := FindUser{}
	err := conn.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return FindUser{},false
	}
	return result,true
}

//Register 用户注册，注册成功返回uid
func Register(user InsertUser) error{
	_, err := conn.InsertOne(context.TODO(), user)
	if err != nil {
		return globalError.DBError("MongoDB异常,",globalError.MongoDBException)
	}
	return nil
}

//UpdateUser 领取奖励更新用户信息
func UpdateUser(all map[uint32]uint64,name string) error{
	_, err := conn.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{
		{"$set", bson.D{
			{"depot", bson.D{
				{"1001", all[uint32(1001)]},
				{"1002", all[uint32(1002)]},
			}},
		}},
	})
	return err
}