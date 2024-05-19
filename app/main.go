package main

import (
	"web/app/presentation/api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	routers := router.GetRouter()
	routers.Run(":8080")
}