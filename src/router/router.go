package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wennmu/gowe-admin.git/pkg/e"
	"github.com/wennmu/gowe-admin.git/src/controllor"
	"github.com/wennmu/gowe-admin.git/src/middleware"
)

func Init() *gin.Engine {
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob("src/view/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	//登录
	router.POST("/login", e.ErrorWrapper(controllor.Login))

	//控制台路由组
	admin := router.Group("/admin")
	{
		page := admin.Group("/page")

		page.Use(middleware.Request())

		page.Use(middleware.Check())
		{
			page.GET("/index", controllor.Index)
		}

		//api := admin.Group("/api")
		//{
		//	api.GET("/index", e.ErrorWrapper(controllor.xxxx))
		//}
	}

	return router
}
