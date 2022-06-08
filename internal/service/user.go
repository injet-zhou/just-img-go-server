package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
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
	var err error
	user, err = user.GetByUsername(global.DBEngine)
	if err != nil {
		return nil, err
	}
	return user, nil
}
