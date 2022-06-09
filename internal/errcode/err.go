package errcode

import (
	"fmt"
)

type Error struct {
	code          int
	encryptedCode string
	msg           string
	extra         []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.msg)
}
