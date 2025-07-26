package routers

import (
	"backend/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// swagger使用
	// 测试qq登录接口

	//如有需求在这里读取json错误码文件

	routerGroup := router.Group("/api")

	TestRouter(routerGroup)

	return router
}
