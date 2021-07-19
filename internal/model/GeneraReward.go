package model

//GeneraReward 通用奖励信息
type GeneraReward struct {
	Code    int32
	Msg     string
	Changes map[uint32]uint64
	Balance map[uint32]uint64
	Counter map[uint32]uint64
}
