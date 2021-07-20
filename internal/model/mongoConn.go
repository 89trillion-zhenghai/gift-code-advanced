package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const MongoUrl = "mongodb://localhost:27017"
var mgoCli *mongo.Client

func InitMongoDB() error{
	var err error
	clientOptions := options.Client().ApplyURI(MongoUrl)
	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		InitMongoDB()
	}
	return mgoCli
}

func GetMgoCol(db string,col string) *mongo.Collection {
	return GetMgoCli().Database(db).Collection(col)
}
