package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FindUser struct {
	Uid   primitive.ObjectID `_id` //mongo中自增主键
	Name  string             //用户名
	Depot map[uint32]uint64  //背包	1001:金币		1002:钻石
}
type InsertUser struct {
	Name  string            //用户名
	Depot map[uint32]uint64 //背包	1001:金币		1002:钻石
}
