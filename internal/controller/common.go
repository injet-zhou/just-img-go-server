package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
)

type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) Error(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}

func Error(ctx *gin.Context, statusCode int, msg string) {
	res := &Response{
		Code:       statusCode,
		Msg:        msg,
		StatusCode: statusCode,
	}
	res.Error(ctx)
}

func Success(ctx *gin.Context, msg string, content interface{}) {

	res := &Response{
		Code:       200,
		Msg:        msg,
		StatusCode: 200,
		Data:       content,
	}
	res.Error(ctx)
}

func ErrorResponse(ctx *gin.Context, err *errcode.Error, msg string) {
	switch err.Code {
	case errcode.ErrWrongUsername:
		Error(ctx, 400, msg)
	case errcode.ErrWrongPassword:
		Error(ctx, 400, msg)
	case errcode.ErrLoginFailTooManyTimes:
		Error(ctx, 403, msg)
	case errcode.ErrUserNameOrEmailRequired:
		Error(ctx, 400, msg)
	case errcode.ErrUserNotExist:
		Error(ctx, 404, msg)
	case errcode.ErrTokenUnauthorized:
		Error(ctx, 401, msg)
	case errcode.ErrTokenExpired:
		Error(ctx, 401, msg)
	case errcode.ErrPasswordRequired:
		Error(ctx, 400, msg)
	case errcode.ErrLoginNameExist:
		Error(ctx, 400, msg)
	default:
		Error(ctx, 500, msg)
	}
}
