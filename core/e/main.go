package e

import (
	"fmt"
	"io"
	"runtime"
)

var (
	ErrNotWx         = Error{code: 101, desc: "请使用微信打开"}
	ErrNotLogin      = Error{code: 102, desc: "用户未登录"}
	ErrNoPermission  = Error{code: 103, desc: "用户无权限"}
	ErrRegistFail    = Error{code: 104, desc: "绑定失败"}
	ErrInvalidParams = Error{code: 201, desc: "请求参数错误"}
	ErrQueryFail     = Error{code: 301, desc: "查询失败"}
	ErrNotFound      = Error{code: 302, desc: "找不到记录"}
	ErrCreateFail    = Error{code: 303, desc: "创建失败"}
	ErrUpdateFail    = Error{code: 304, desc: "更新失败"}
	ErrInternal      = Error{code: 500, desc: ""}
)

type Error struct {
	code  uint
	desc  string
	cause error
	stack []uintptr
}

func (e Error) From(err error) Error {
	e.cause = err
	e.stack = callers()
	return e
}

func (e Error) Code() uint {
	return e.code
}

func (e Error) Error() string {
	if e.cause != nil {
		return e.desc + ":" + e.cause.Error()
	}
	return e.desc
}

func (e Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Error())
			io.WriteString(s, "\n")
			frames := runtime.CallersFrames(e.stack)
			for frame, more := frames.Next(); more; frame, more = frames.Next() {
				fmt.Fprintf(s, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
			}
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, e.Error())
	}
}

func callers() []uintptr {
	pcs := make([]uintptr, 10)
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}
