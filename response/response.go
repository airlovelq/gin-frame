package response

// 错误处理的结构体
type Response struct {
	StatusCode int         `json:"-"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
