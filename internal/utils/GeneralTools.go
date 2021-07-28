package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)
//SendComplexGetRequest 通过调考核题三的接口完成对礼品内容的获取是否合法
func SendComplexGetRequest(gc string, name string) map[string]map[string]interface{} {
	params := url.Values{"giftCode": {gc}, "uuid": {name}}
	resp, _ := http.PostForm("http://127.0.0.1:8080/redeemGift", params)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resMap := make(map[string]map[string]interface{})
	json.Unmarshal(body, &resMap)
	return resMap
}
//MapToAnther 转其他类型的map
func MapToAnther(params map[string]string) map[uint32]uint64 {
	res := make(map[uint32]uint64)
	for k, v := range params {
		key, err := strconv.Atoi(k)
		value, err := strconv.Atoi(v)
		if err != nil {
		}
		res[uint32(key)] = uint64(value)
	}
	return res
}

//DepotAdd 将两个map的相同key的value值进行求和
func DepotAdd(m1 map[uint32]uint64, m2 map[uint32]uint64) map[uint32]uint64 {
	sum := make(map[uint32]uint64)
	for k, _ := range m1 {
		if value,ok:=m2[k]; ok{
			sum[k] = m1[k] + value
		}else{
			sum[k] = m1[k]
		}
	}
	return sum
}
