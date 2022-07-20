package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/dao"
	"github.com/injet-zhou/just-img-go-server/internal/service"
)

func ImageListController(ctx *gin.Context) {
	req := new(service.ImagesRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		Error(ctx, 400, err.Error())
		return
	}
	user := ctx.Keys["User"].(*dao.TokenUser)
	if !user.IsAdmin {
		req.Username = user.Username
		req.GroupName = ""
	}
	images, err := service.ImageList(req)
	if err != nil {
		Error(ctx, 500, err.Error())
		return
	}
	Success(ctx, "success", images)
}
