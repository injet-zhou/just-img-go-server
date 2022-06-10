package errcode

import (
	"fmt"
)

type Error struct {
	code  int
	msg   string
	extra []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.msg)
}

func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
