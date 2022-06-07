package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/tool"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Login(ctx *gin.Context) (*entity.User, error) {
	user := &entity.User{}
	req := &LoginRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	req = tool.TrimFields(req).(*LoginRequest)
	if req.Username == "" && req.Email == "" {
		return nil, fmt.Errorf("username or email is required")
	}

	return user, nil
}
