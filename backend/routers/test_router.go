package routers

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func TestRouter(router *gin.RouterGroup) {
	testApi := api.ApiGroupApp.TestApi

	router.GET("/test", testApi.TestView)
}
