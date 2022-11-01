package utils

import "encoding/json"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 自定义返回相应信息
func (res *Response) WithMessage(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

func (res *Response) ToString() string {
	err := &struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}

func response(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
