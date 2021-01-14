package errno

import (
	"encoding/json"
)

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithData 设置成功时返回的数据
	WithData(data interface{}) Error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewError(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (e *err) i() {}

func (e *err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
