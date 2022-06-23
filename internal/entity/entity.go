package entity

import (
	"errors"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	if db != nil {
		return db.AutoMigrate(
			&User{},
			&UserGroup{},
			&Image{},
		)
	}
	return errors.New("db is nil")
}

func CreateAdminUser() error {
	log := logger.Default()
	hasAdminGroup, err := isAdminGroupExist()
	if err != nil {
		log.Error("isAdminGroupExist error: ", zap.String("err", err.Error()))
		return err
	}
	if hasAdminGroup {
		return nil
	}
	accountCfg := config.GetAccountCfg()
	user := &User{}
	if accountCfg == nil {
		log.Warn("*****[UNSAFE]****** account config is nil")
		user.Username = "just_admin"
		user.Password = "just_123987654"
	}
	db := global.DBEngine
	transErr := db.Transaction(func(tx *gorm.DB) error {
		userGroup := &UserGroup{
			Name:    "admin",
			IsAdmin: true,
		}
		if err := tx.Create(userGroup).Error; err != nil {
			return err
		}
		user.GroupId = userGroup.ID
		if accountCfg != nil {
			user.Username = accountCfg.Username
			user.Password = accountCfg.Password
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	if transErr != nil {
		return transErr
	}
	return nil
}

func isAdminGroupExist() (bool, error) {
	userGroup := &UserGroup{}
	db := global.DBEngine
	err := db.Where("is_admin = ?", true).First(userGroup).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
