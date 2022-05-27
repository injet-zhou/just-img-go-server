package controller

import "github.com/gin-gonic/gin"

type Response struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	StatusCode int
	Content    *interface{} `json:"content,omitempty"`
}

func (r *Response) Error(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}

func ErrorResponse(ctx *gin.Context, statusCode int, msg string) {
	res := &Response{
		Code:       statusCode,
		Msg:        msg,
		StatusCode: statusCode,
	}
	res.Error(ctx)
}

func Success(ctx *gin.Context, msg string, content *interface{}) {
	res := &Response{
		Code:       200,
		Msg:        msg,
		StatusCode: 200,
		Content:    content,
	}
	res.Error(ctx)
}
