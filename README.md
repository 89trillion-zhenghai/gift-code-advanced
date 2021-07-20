## 项目简介

​	1、基于考题三进行开发两个新接口以及对考题三点接口进行修改。

​	2、项目需要实现保存用户注册信息和校验登陆信息。登陆成功则返回用户有关的数据包括UID、金币数、钻石数。

## 快速上手

​	1、本项目由Go语言开发，数据库采用redis+mongo。需要配置Go开发环境，以及安装redis数据库和mongodb数据库

​	2、进入项目根目录，通过命令行命令启动项目

​			1、go build main.go		---编译

​			2、./main 						  ---启动项目

## 目录结构

![wecom-temp-e850fcff31962eecd615b34808f8d3ad](/var/folders/yh/qxwd_mm96jd6l4_hbk8q_cmr0000gp/T/com.tencent.WeWorkMac/wecom-temp-e850fcff31962eecd615b34808f8d3ad.png)

## 代码逻辑分层

| 层      | 文件夹                     | 功能                         | 调用关系                 | 其他说明     |
| ------- | -------------------------- | ---------------------------- | ------------------------ | ------------ |
| 应用层  | /app/http/httpServer.go    | 启动服务器                   | 调用路由层               | 不可同层调用 |
| 路由层  | /internal/router/router.go | 路由转发                     | 被应用层调用，调用控制层 | 不可同层调用 |
| 控制层  | /internal/ctrl             | 请求参数处理，调用逻辑层处理 | 被路由层调用，调用逻辑层 | 不可同层调用 |
| 逻辑层  | /internal/service          | 处理具体业务                 | 被控制层调用             | 不可同层调用 |
| 工具层  | /internal/utils            | 通用工具                     | 被其他层调用             | 可同层调用   |
| model层 | /internal/model            | 数据模型、数据库操作         | 被控制层和逻辑层调用     | 可同层调用   |



## 功能介绍

​	1、用户登陆时需要判断用户是否是新用户，如果是新用户则注册新用户，并且为新用户生成一个UID作为唯一标识。如果不是新用户，则将用户数据返回给用户，包括唯一UID、金币数、钻石数。	

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

请求地址：http://127.0.0.1:8080/createAndGetGiftCode

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
   "giftCode": "35CY3RAS"
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
    "message": {
        "Uid": "60f6a6e273f8665b9d01dbe1",
        "Name": "smallbai",
        "Depot": {
            "1001": 10,
            "1002": 5
        }
    }
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
    "message": "CMgBEgzpooblj5bmiJDlip8aBQjpBxAKGgUI6gcQBSIFCOkHEAAiBQjqBxAAKgUI6QcQCioFCOoHEAUyDOaJqeWxleWtl+autQ=="
}
```

#### 响应状态码

| 状态码 | 说明              |
| ------ | ----------------- |
| 无     | 成功              |
| 1001   | 礼品码不可领取    |
| 1002   | 未知错误          |
| 1003   | mongodb数据库异常 |
| 1004   | 参数为空          |

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
