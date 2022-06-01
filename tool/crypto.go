package tool

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(msg string) string {
	data := []byte(msg)
	ctx := md5.New()
	ctx.Write(data)
	cipherStr := ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
