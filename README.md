## 项目简介

​	1、基于考题三进行开发两个新接口以及对考题三点接口进行修改。

​	2、项目需要实现保存用户注册信息和校验登陆信息。登陆成功则返回用户有关的数据包括UID、金币数、钻石数。

## 快速上手

​	1、本项目由Go语言开发，数据库采用redis+mongo。需要配置Go开发环境，以及安装redis数据库和mongodb数据库

​	2、进入项目根目录，通过命令行命令启动项目

​			1、go build main.go		---编译

​			2、./main 						  ---启动项目

## 目录结构

```tree
├── README.md
├── __pycache__
│   └── locust.cpython-39.pyc
├── app
│   ├── http
│   │   └── httpServer.go
│   ├── main
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   ├── giftController.go
│   │   └── userCtrl.go
│   ├── globalError
│   │   ├── error.go
│   │   └── errorHandler.go
│   ├── handler
│   │   ├── giftHandler.go
│   │   ├── handler_test.go
│   │   └── userHandler.go
│   ├── model
│   │   ├── GeneraReward.go
│   │   ├── GifCodetInfo.go
│   │   ├── mongoConn.go
│   │   ├── mongoOpt.go
│   │   ├── redisConn.go
│   │   ├── redisOpt.go
│   │   └── user.go
│   ├── router
│   │   └── route.go
│   ├── service
│   │   ├── giftService.go
│   │   └── userService.go
│   ├── utils
│   │   ├── GeneralTools.go
│   │   └── GeneralUtils.go
│   └── verify
│       └── paramter.go
├── locust.py
├── report.html
└── response
    ├── GeneralReward.pb.go
    └── GeneralReward.proto
```



## 代码逻辑分层

| 层          | 文件夹                     | 功能                               | 调用关系                    |
| ----------- | -------------------------- | ---------------------------------- | --------------------------- |
| 应用层      | /app/http/httpServer.go    | 启动服务器                         | 调用路由层                  |
| 路由层      | /internal/router/router.go | 路由转发                           | 被应用层调用，调用控制层    |
| 控制层      | /internal/ctrl             | 请求参数处理，调用handler层处理    | 被路由层调用，调用handler层 |
| handler     | /internal/handler          | 通用业务处理，具体业务调业务层处理 | 调用业务层                  |
| 业务层      | /internal/service          | 处理具体业务                       | 被handler层调用             |
| 工具层      | /internal/utils            | 通用工具                           | 被其他层调用                |
| model层     | /internal/model            | 数据模型、数据库操作               | 被控制层和业务层调用        |
| verify      | /internal/verify           | 数据校验                           | 被其他层调用                |
| globalError | /internal/globalError      | 全局错误处理                       | 被其他层调用                |



## 功能介绍

​	1、用户登陆时需要判断用户是否是新用户，如果是新用户则注册新用户，并且为新用户生成一个UID作为唯一标识。如果不是新用户，则将用户数据返回给用户，包括唯一UID、name、金币数、钻石数。	

​	2、对考题三的验证礼品码接口修改，用户的奖励信息储存到mongo数据库，接口返回值应该是protobuf对象的[]byte，客户端需要一个protobuf解析函数对返回到[]byte进行解析并将解析后的内容展示给用户

## 存储设计

用户信息

| 内容         | 数据库  | Key   |
| ------------ | ------- | ----- |
| 用户唯一标识 | MongoDB | Uid   |
| 用户名       | MongoDB | Name  |
| 用户仓库     | MongoDB | Depot |

## 接口设计

#### 1、创建礼品码接口，需要考核题三的支持

请求方法：POST

请求地址：http://127.0.0.1:8000/createAndGetGiftCode

请求参数: form-data

```json
userName :	admin
description : 十周年活动奖励
giftType : 2
validity : 10m
availableTimes : 800
giftDetail : {"1001":"10","1002":"5"}
```

请求响应：

```json
{
    "data": "JF362262",
    "message": "success",
    "status": 200
}
```

#### 2、登陆/注册接口

请求方法：POST

请求地址：http://127.0.0.1:8000/login

请求参数：form-data

```json
name ： smallbai
```

请求响应：

```json
{
    "data": {
        "Uid": "61012eda7f46c0dc0f72f5ad",
        "Name": "testName02",
        "Depot": {
            "1001": 15,
            "1002": 10
        }
    },
    "message": "success",
    "status": 200
}
```

#### 3、兑换礼品接口

请求方法：POST

请求地址：http://127.0.0.1:8000/redeemGift

请求参数：form-data

```json
name : smallbai
giftCode : 35CY3RAS(应调用接口一获取有效礼品码)
```

​	请求响应

```json
{
    "data": "CMgBEgzpooblj5bmiJDlip8aBQjpBxAPGgUI6gcQCiIFCOkHEAAiBQjqBxAAKgUI6QcQDyoFCOoHEAoyDOaJqeWxleWtl+autQ==",
    "message": "success",
    "status": 200
}
```

#### 4、管理员查询礼品码接口

#### 请求方法

http GET

#### 接口地址

http://127.0.0.1:8000/getGiftDetail

#### 请求参数

```
giftCode=6502M6S6
```

#### 请求响应

```
{
    "data": {
        "AvailableDetail": {
            "smallbai": "2021-07-27 21:05:13",
            "yangzhenghai": "2021-07-27 21:05:28"
        },
        "AvailableTime": "2",
        "AvailableTimes": "20",
        "CreateTime": "2021-07-27 21:02:31",
        "CreateUser": "admin",
        "Description": "十周年纪念",
        "GiftCode": "6502M6S6",
        "GiftDetail": "{\"1001\":\"2\",\"1003\":\"3\"}",
        "GiftType": "2",
        "Validity": "1627391551"
    },
    "status": 200
}
```

#### 响应状态码

| 状态码 | 说明                     |
| ------ | ------------------------ |
| 200    | 成功                     |
| 1001   | 礼品码已过期             |
| 1002   | 该用户已经领取过礼品码了 |
| 1003   | 礼品码不存在/错误        |
| 1004   | 礼品码已失效             |
| 1005   | 礼品被领取完毕           |
| 1006   | 服务器异常               |
| 1007   | mongodb数据库异常        |
| 1008   | redis数据库异常          |
| 1009   | 参数为空                 |
| 1010   | 参数不合法               |
| 1011   | 用户不存在               |
| 1012   | 用户注册                 |

## 第三方库

### gin

```
go语言的web框架
https://github.com/gin-gonic/gin
```

### go-redis

```
go语言连接操作redis数据库
https://github.com/go-redis
```

### Mongoldb go

```
go语言mongdb官方驱动，连接操作mongodb数据库
https://go.mongodb.org/mongo-driver/mongo
```

### protobuf

```
包含go语言处理proto数据的函数
http://github.com/golang/protobuf/proto
```

## 

## todo

将代码进一步分层

## 流程图

![未命名文件 (7)](https://user-images.githubusercontent.com/86946999/125749011-89938e98-5f72-442f-997b-41d4bf19f8db.jpg)
