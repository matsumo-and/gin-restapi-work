package router

import (
	"web/app/presentation/api/v1/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", controller.Index)
	return r
}