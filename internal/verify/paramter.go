package verify

import "strconv"

//ParamIsNotEmpty 校验参数是否为空 为空返回false
func ParamIsNotEmpty(params ...string) bool {
	for _, param := range params {
		if len(param) == 0{
			return false
		}
	}
	return true
}

//IsDigit 校验参数是否为数字
func IsDigit(str string) bool{
	_, err := strconv.Atoi(str)
	return err == nil
}
