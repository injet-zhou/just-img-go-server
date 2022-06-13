package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/app"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
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
		newErr := err.(*errcode.Error)
		if newErr.Code == errcode.ErrWrongUsername || newErr.Code == errcode.ErrWrongPassword {
			ErrorResponse(ctx, 400, newErr.Msg)
			return
		}
		if newErr.Code == errcode.ErrLoginFailTooManyTimes {
			ErrorResponse(ctx, 403, newErr.Msg)
			return
		}
		if newErr.Code == errcode.ErrUserNameOrEmailRequired {
			ErrorResponse(ctx, 400, newErr.Msg)
			return
		}
		if newErr.Code == errcode.ErrUserNotExist {
			ErrorResponse(ctx, 404, newErr.Msg)
			return
		}
		ErrorResponse(ctx, 500, newErr.Msg)
		return
	}
	safeUser := user.SafeInfo()
	token, err := app.GenToken(user)
	if err != nil {
		ErrorResponse(ctx, 500, err.Error())
		return
	}
	safeUser.Token = token
	Success(ctx, "login success", safeUser)
}
