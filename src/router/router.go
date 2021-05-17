package router

import (
	"../controllor"
)

//import "../../pkg/e"
import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob("src/view/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//控制台路由组
	admin := router.Group("/admin")
	{
		page := admin.Group("/page")
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
