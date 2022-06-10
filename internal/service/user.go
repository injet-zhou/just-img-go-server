package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Login(ctx *gin.Context, req *LoginRequest) (*entity.User, error) {
	user := &entity.User{}
	if req.Username == "" && req.Email == "" {
		return nil, fmt.Errorf("username or email is required")
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	var err error
	user, err = user.GetByLoginName(global.DBEngine)
	if err != nil {
		return nil, errcode.NewError(errcode.ErrWrongUsername, "账号不存在")
	}
	return user, nil
}

//func CanILogin(ctx *gin.Context, user *entity.User) bool {
//
//}
