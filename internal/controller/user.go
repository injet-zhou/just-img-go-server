package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/service"
	"github.com/injet-zhou/just-img-go-server/tool"
)

func Login(ctx *gin.Context) {
	req := service.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, 400, err.Error())
		return
	}
	req = tool.TrimFields(&req).(service.LoginRequest)
}
