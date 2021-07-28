package utils

import (
	"math/rand"
	"strings"
	"time"
)


//GetGiftCode 随机获得8位礼品码 由大写字母和数字组成
func GetGiftCode() string {
	str := make([]string,6)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		bt := rand.Intn(10)
		if bt > 5 {
			bt = rand.Intn(26)+65
		}else{
			bt = rand.Intn(9)+48
		}
		str = append(str, string(rune(bt)))
	}
	return strings.Join(str,"")
}
