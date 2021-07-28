package globalError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	GiftCodeExpired  	= 1001		//礼品码已过期
	GiftCodeReceived 	= 1002		//该用户已经领取过礼品码了
	GiftCodeNotExist 	= 1003		//礼品码不存在/错误
	GiftCodeIsInvalid	= 1004		//礼品码已失效
	GiftIsOver		 	= 1005		//礼品被领取完毕
	ServerException    	= 1006 		//服务器异常
	MongoDBException   	= 1007 		//mongodb数据库异常
	RedisException		= 1008		//redis数据库异常
	ParamIsEmpty	 	= 1009		//参数为空
	ParamIsIllegal		= 1010		//参数不合法
	UserNotExit			= 1011		//用户不存在
	UserRegister		= 1012		//用户注册
)

type GlobalHandler func(c *gin.Context) (interface{}, error)


//ErrorHandler 全局异常处理
func ErrorHandler(handler GlobalHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			globalError := err.(GlobalError)
			c.JSON(globalError.Status, globalError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":http.StatusOK,
			"data":data,
			"message":"success",
		})
	}
}

// ParamError 参数异常
func ParamError(message string, code int) GlobalError {
	return GlobalError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
	}
}

// ServerError 服务器内部异常
func ServerError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    ServerException,
		Message: message,
	}
}
//DBError 数据库异常
func DBError(message string,code int) GlobalError {
	return GlobalError{
		Status: http.StatusInternalServerError,
		Code: code,
		Message: message,
	}
}


//GiftCodeError 礼品码异常
func GiftCodeError(message string, code int) GlobalError {
	return GlobalError{
		Status: http.StatusBadRequest,
		Code: code,
		Message: message,
	}
}
//UserError 用户不存在
func UserError(message string) GlobalError {
	return GlobalError{
		Status: http.StatusBadRequest,
		Code: UserNotExit,
		Message: message,
	}
}

//Register 用户注册
func Register(name string) GlobalError {
	return GlobalError{
		Status: http.StatusOK,
		Code: UserRegister,
		Message: "欢迎新用户注册，您的通信证为："+name,
	}
}