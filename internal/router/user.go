package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/controller"
	"github.com/injet-zhou/just-img-go-server/internal/middleware"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
		user.Use(middleware.AuthMiddleware())
		user.POST("/list", controller.UserListController)
	}
}
