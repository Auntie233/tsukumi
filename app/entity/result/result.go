package result

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeSuccess = 0
)

func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

func (res *Result) SetMessage(message string) *Result {
	res.Message = message
	return res
}

func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
