package errcode

import (
	"fmt"
	"github.com/injet-zhou/just-img-go-server/config"
	"os"
)

type Error struct {
	code          int
	encryptedCode string
	msg           string
	extra         []string
}

func (e *Error) Error() string {
	if os.Getenv(config.ENVkEY) == config.PROD {
		return fmt.Sprintf("错误码：%s，错误信息：%s", e.encryptedCode, e.msg)
	}
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.msg)
}
