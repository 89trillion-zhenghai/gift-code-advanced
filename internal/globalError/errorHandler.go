package globalError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	NotConvertible	   = 1001 //礼品码不可领取
	ServerException    = 1002 //未知错误
	MongoDBException   = 1003 //mongodb数据库异常
		ParamIsEmpty       = 1004 //参数为空
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
		c.JSON(http.StatusOK, data)
	}
}

// ExpressionError 参数异常
func ExpressionError(message string, code int) GlobalError {
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
//MongoDBError mongodb数据库异常
func MongoDBError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    MongoDBException,
		Message: message,
	}
}
//Param 参数为空
func Param(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    ParamIsEmpty,
		Message: message,
	}
}

//GiftNotConvertible 礼品不可以领取
func GiftNotConvertible(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    NotConvertible,
		Message: message,
	}
}