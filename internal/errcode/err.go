package errcode

import (
	"fmt"
)

type Error struct {
	Code  int
	Msg   string
	Extra []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code, e.Msg)
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
