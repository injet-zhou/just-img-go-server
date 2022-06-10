package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/pb"
	"github.com/injet-zhou/just-img-go-server/internal/service"
	"github.com/injet-zhou/just-img-go-server/tool"
)

func Login(ctx *gin.Context) {
	req := &pb.LoginRequest{}
	if err := ctx.Bind(req); err != nil {
		ErrorResponse(ctx, 400, err.Error())
		return
	}
	loginReq := &service.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	loginReq = tool.TrimFields(loginReq).(*service.LoginRequest)
	user, err := service.Login(ctx, loginReq)
	if err != nil {
		ErrorResponse(ctx, 500, err.Error())
		return
	}
	safeUser := user.SafeInfo()
	Success(ctx, "login success", safeUser)
}
