package e

import "errors"

type Error struct {
	error
	code uint
}

var (
	// 服务级错误
	ErrInternal       = Error{errors.New("未知错误"), 10101}
	ErrTokenRefresher = Error{errors.New("Token更新故障，请尽快排查"), 10102}

	//请求级错误
	ErrParam     = Error{errors.New("请求参数有误"), 20101}
	ErrSignature = Error{errors.New("签名信息有误"), 20102}
)

func (e Error) Code() uint {
	return e.code
}
