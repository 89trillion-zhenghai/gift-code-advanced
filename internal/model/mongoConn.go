package model

import (
	"context"
	"gift-code-Two/internal/globalError"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Collection

func InitMongoDB() error{
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	mgoCli, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return globalError.DBError("MongoDB连接异常",globalError.MongoDBException)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		return globalError.DBError("MongoDB连接异常",globalError.MongoDBException)
	}
	conn = mgoCli.Database("user").Collection("depot")
	return nil
}

