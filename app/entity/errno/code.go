package errno

var (
	ErrServer    = NewError(10001, "服务异常，请联系管理员")
	ErrParam     = NewError(10002, "参数有误")
	ErrSignParam = NewError(10003, "签名参数有误")
)
