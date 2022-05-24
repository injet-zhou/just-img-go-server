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
