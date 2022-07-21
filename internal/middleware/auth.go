package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/injet-zhou/just-img-go-server/internal/app"
	"github.com/injet-zhou/just-img-go-server/internal/dao"
	"github.com/injet-zhou/just-img-go-server/internal/errcode"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"go.uber.org/zap"
)

func AuthMiddleware() gin.HandlerFunc {
	log := logger.Default()
	return func(c *gin.Context) {
		var token string
		token = c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code":    errcode.ErrTokenUnauthorized,
				"message": "token is empty",
			})
			return
		}
		claims, err := app.ParseToken(token)
		if err != nil {
			log.Error("parse token error", zap.String("err", err.Error()))
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				c.JSON(401, gin.H{
					"code":    errcode.ErrTokenExpired,
					"message": "token is expired",
				})
				return
			default:
				c.JSON(401, gin.H{
					"code":    errcode.ErrTokenUnauthorized,
					"message": "token is invalid",
				})
				return
			}
		}
		c.Set("UserId", claims.UserId)
		c.Set("UserName", claims.Username)
		user, findUserErr := dao.GetUser(claims.UserId)
		if findUserErr != nil {
			log.Error("get user error", zap.String("err", findUserErr.Error()))
			c.JSON(401, gin.H{
				"code":    errcode.ErrTokenUnauthorized,
				"message": "token is invalid",
			})
			return
		}
		c.Set("User", user)
		c.Next()
	}
}
