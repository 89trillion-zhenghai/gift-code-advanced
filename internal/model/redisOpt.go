package model

import (
	"gift-code-Two/internal/globalError"
	"log"
	"time"
)

//GiftIsExit 判断礼品码是否已存在，存在返回false，不存在返回ture
func GiftIsExit(code string) bool {
	result, _ := Rc.Keys("MAIN_" + code).Result()
	return len(result) == 0
}

//SetGift 保存礼品码信息,将礼品码主体信息保存到hash
func SetGift(gift Gift) error {
	//礼品码主体信息保存到hash里
	gMap := gift.BeanToMap()
	_, err := Rc.HMSet("MAIN_"+gift.GiftCode, gMap).Result()
	return err
}

//GetGift 查询礼品码主体信息
func GetGift(code string) (map[string]string,error){
	resMap, err := Rc.HGetAll("MAIN_" + code).Result()
	if err != nil {
		return nil,err
	}
	return resMap,nil
}

//SetAvailableTime 领取次数保存到string
func SetAvailableTime(code string) error {
	_, err := Rc.Set("AVAILABLE_"+code,"0",0).Result()
	return err
}

//GetAvailableTime 查询领取次数
func GetAvailableTime(code string) (string, error){
	result, err := Rc.Get("AVAILABLE_" + code).Result()
	return result,err
}

//SetAvailableDetail 领取列表信息保存到hash
func SetAvailableDetail(code string) error {
	_, err := Rc.HMSet("DETAIL_"+code, nil).Result()
	return err
}

func GetAvailableDetail(code string) (map[string]string,error) {
	result, err := Rc.HGetAll("DETAIL_" + code).Result()
	return result,err
}

//GiftIsAvailed 判断礼品是否被领取过 领取过返回false
func GiftIsAvailed(code string) bool {
	result, _ := Rc.Keys("DETAIL_" + code).Result()
	return len(result) == 0
}

//UserIsAvailed 判断用户是否领取过 领取过返回false
func UserIsAvailed(code string,name string) bool {
	result, _ := Rc.HGet("DETAIL_"+code, name).Result()
	return len(result) == 0
}

//IncrAvailableAndAppendUser 开启事务，领取次数+1和领取列表更新
func IncrAvailableAndAppendUser(code string,name string)  error{
	now := time.Now().Format("2006-01-02 15:04:05")
	pipe := Rc.TxPipeline()
	//领取次数+1
	pipe.Incr("AVAILABLE_" + code)
	//领取列表填充
	pipe.HSet("DETAIL_"+code, name, now)
	_, err := pipe.Exec()
	if err != nil{
		//取消提交
		pipe.Discard()
		return globalError.DBError("redis数据库发送错误！！",globalError.RedisException)
	}
	return nil
}

func Rollback(code string , user FindUser) {
	//回退操作需要把领取列表回退和领取次数-1
	pipe := Rc.TxPipeline()
	pipe.Decr("AVAILABLE_" + code)
	pipe.HDel("DETAIL_"+code, user.Name)
	_, err := pipe.Exec()
	if err != nil {
		//取消提交
		pipe.Discard()
		log.Fatalln(err.Error())
	}
	err = UpdateUser(user.Depot, user.Name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
