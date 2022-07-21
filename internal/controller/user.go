package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/app"
	"github.com/injet-zhou/just-img-go-server/internal/dao"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
	"github.com/injet-zhou/just-img-go-server/internal/pb"
	"github.com/injet-zhou/just-img-go-server/internal/service"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"github.com/injet-zhou/just-img-go-server/tool"
	"go.uber.org/zap"
)

func bindParams(c *gin.Context, params interface{}, module string) error {
	log := logger.Default()
	if err := c.Bind(params); err != nil {
		log.Error("bind params error", zap.String("err", err.Error()), zap.String("module", module))
		Error(c, 400, err.Error())
		return err
	}
	return nil
}

func Login(ctx *gin.Context) {
	req := &pb.LoginRequest{}
	log := logger.Default()
	if err := bindParams(ctx, req, "login"); err != nil {
		return
	}
	loginReq := &service.AuthRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	loginReq = tool.TrimFields(loginReq).(*service.AuthRequest)
	user, err := service.Login(ctx, loginReq)
	if err != nil {
		newErr := err.(*errcode.Error)
		ErrorResponse(ctx, newErr, newErr.Msg)
		return
	}
	safeUser := user.SafeInfo()
	token, err := app.GenToken(user)
	if err != nil {
		log.Error("generate token error", zap.String("err", err.Error()))
		Error(ctx, 500, err.Error())
		return
	}
	safeUser.Token = token
	Success(ctx, "login success", safeUser)
}

func Register(c *gin.Context) {
	req := &pb.RegisterRequest{}
	if err := bindParams(c, req, "register"); err != nil {
		return
	}
	registerReq := &service.AuthRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	registerReq = tool.TrimFields(registerReq).(*service.AuthRequest)
	user, err := service.Register(c, registerReq)
	if err != nil {
		newErr := err.(*errcode.Error)
		ErrorResponse(c, newErr, newErr.Msg)
		return
	}
	safeUser := user.SafeInfo()
	token, err := app.GenToken(user)
	if err != nil {
		logger.Default().Error("generate token error", zap.String("err", err.Error()))
		Error(c, 500, err.Error())
		return
	}
	safeUser.Token = token
	Success(c, "register success", safeUser)
}

func UserListController(ctx *gin.Context) {
	user := ctx.Keys["User"].(*dao.TokenUser)
	if !user.IsAdmin {
		Error(ctx, 403, "forbidden")
		return
	}
}
