package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/dao"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
	"github.com/injet-zhou/just-img-go-server/pkg"
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
	user.Password = req.Password
	userEntity, err := user.GetByLoginName(global.DBEngine)
	isLoginNameExist := false
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("get user by login name error", zap.String("err", err.Error()))
			return nil, errcode.NewError(errcode.DBErr, err.Error())
		}
	}
	if userEntity != nil && userEntity.ID > 0 {
		isLoginNameExist = true
	}
	if isLoginNameExist {
		return nil, errcode.NewError(errcode.ErrLoginNameExist, "用户名已存在")
	}
	defaultGroup, groupErr := dao.DefaultUserGroup(global.DBEngine)
	if groupErr != nil {
		log.Error("get default user group error", zap.String("err", groupErr.Error()))
		return nil, errcode.NewError(errcode.DBErr, groupErr.Error())
	}
	user.GroupId = defaultGroup.ID
	createErr := user.Create(global.DBEngine)
	if createErr != nil {
		log.Error("create user error", zap.String("err", createErr.Error()))
		return nil, errcode.NewError(errcode.DBErr, createErr.Error())
	}
	return user, nil
}

// PasswordValidate 检查密码
func PasswordValidate(password, dbPassword, UID string) bool {
	psw := tool.MD5(password + UID)
	return strings.Compare(dbPassword, psw) == 0
}

type UsersRequest struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	GroupName      string `json:"groupName"`
	CreatedAtStart string `json:"createdAtStart"`
	CreatedAtEnd   string `json:"createdAtEnd"`
	UID            string `json:"uid"`
	Nickname       string `json:"nickname"`
}

func UserList(req *UsersRequest) (*pkg.Pagination, error) {
	if req == nil {
		return nil, errcode.NewError(errcode.ErrInvalidParams, "invalid params")
	}
	db := global.DBEngine.Model(&entity.User{})
	db = db.Select("user.id,user.username,user.email,user.group_id,user.created_at,user.updated_at,user.uid,user.nickname,user.avatar,user_group.name as group_name")
	db = db.Joins("left join user_group on user.group_id = user_group.id")
	likeWrapper := tool.LikeQueryWrapper
	if req.Username != "" {
		db = db.Where("user.username like ?", likeWrapper(req.Username))
	}
	if req.Email != "" {
		db = db.Where("user.email like ?", likeWrapper(req.Email))
	}
	if req.GroupName != "" {
		db = db.Where("user_group.name like ?", likeWrapper(req.GroupName))
	}
	dateStart := tool.DateStr2Timestamp(req.CreatedAtStart)
	dateEnd := tool.DateStr2Timestamp(req.CreatedAtEnd)
	if dateStart != 0 && dateEnd != 0 && dateStart > dateEnd {
		return nil, errcode.NewError(errcode.ErrInvalidParams, "date start must less than date end")
	}
	if req.CreatedAtStart != "" {
		db = db.Where("user.created_at >= ?", req.CreatedAtStart)
	}
	if req.CreatedAtEnd != "" {
		db = db.Where("user.created_at <= ?", req.CreatedAtEnd)
	}
	if req.UID != "" {
		db = db.Where("user.uid like ?", "'%"+req.UID+"'")
	}
	if req.Nickname != "" {
		db = db.Where("user.nickname like ?", "'%"+req.Nickname+"'")
	}
	paginator := &pkg.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
	}
	db = db.Scopes(pkg.Paginate(paginator, db))
	users, err := dao.GetUsers(db)
	if err != nil {
		return nil, errcode.NewError(errcode.DBErr, err.Error())
	}
	paginator.Rows = users
	return paginator, nil
}
