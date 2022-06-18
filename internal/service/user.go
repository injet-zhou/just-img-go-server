package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"github.com/injet-zhou/just-img-go-server/tool"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Login(ctx *gin.Context, req *AuthRequest) (*entity.User, error) {
	user := &entity.User{}
	if req.Username == "" && req.Email == "" {
		return nil, errcode.NewError(errcode.ErrUserNameOrEmailRequired, "username or email is required")
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	var err error
	user, err = user.GetByLoginName(global.DBEngine)
	log := logger.Default()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NewError(errcode.ErrWrongUsername, "账号不存在")
		}
		log.Error("get user by login name error", zap.String("err", err.Error()))
		return nil, errcode.NewError(errcode.DBErr, err.Error())
	}
	if user.ID > 0 {
		if authFailTimes, err := global.RedisEngine.Get(ctx, fmt.Sprintf("%s:%d", config.USER_SESSION_KEY, user.ID)).Int(); err == nil {
			if authFailTimes > config.MAX_LOGIN_FAIL_COUNT {
				return nil, errcode.NewError(errcode.ErrLoginFailTooManyTimes, "登录失败次数过多")
			}
		} else {
			log.Error("get user login fail times error", zap.String("err", err.Error()))
		}
	}
	if !PasswordValidate(req.Password, user.Password, user.UID) {
		if user.ID > 0 {
			if err := global.RedisEngine.Incr(ctx, fmt.Sprintf("%s:%d", config.USER_SESSION_KEY, user.ID)).Err(); err != nil {
				log.Error("incr user login fail times error", zap.String("err", err.Error()))
			}
		}
		return nil, errcode.NewError(errcode.ErrWrongPassword, "密码错误")
	}
	return user, nil
}

func Register(ctx *gin.Context, req *AuthRequest) (*entity.User, error) {
	user := &entity.User{}
	log := logger.Default()
	if req.Username == "" && req.Email == "" {
		return nil, errcode.NewError(errcode.ErrUserNameOrEmailRequired, "username is required")
	}
	if req.Password == "" {
		return nil, errcode.NewError(errcode.ErrPasswordRequired, "password is required")
	}
	user.Email = req.Email
	user.Username = req.Username
	var err error
	user, err = user.GetByLoginName(global.DBEngine)
	isLoginNameExist := false
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("get user by login name error", zap.String("err", err.Error()))
			return nil, errcode.NewError(errcode.DBErr, err.Error())
		}
		isLoginNameExist = true
	}
	if isLoginNameExist {
		return nil, errcode.NewError(errcode.ErrLoginNameExist, "用户名已存在")
	}
	createErr := user.Create(global.DBEngine)
	if createErr != nil {
		log.Error("create user error", zap.String("err", createErr.Error()))
		return nil, errcode.NewError(errcode.DBErr, createErr.Error())
	}
	return user, nil
}

// PasswordValidate 检查密码
func PasswordValidate(password, dbPassword, UID string) bool {
	return strings.Compare(dbPassword, tool.MD5(password+UID)) == 0
}
