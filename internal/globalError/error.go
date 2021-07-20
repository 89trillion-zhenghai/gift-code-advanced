package globalError

//GlobalError 错误结构体
type GlobalError struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err GlobalError) Error() string {
	return err.Message
}
