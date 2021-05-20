package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wennmu/gowe-admin.git/pkg/e"
	"github.com/wennmu/gowe-admin.git/src/controllor"
	"github.com/wennmu/gowe-admin.git/src/middleware"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob("src/view/*.html")
	router.StaticFS("/static", http.Dir("src/view/static"))
	router.StaticFile("/favicon.ico", "src/view/favicon.ico")

	router.Use(middleware.Request())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	//登录
	router.POST("/login", e.ErrorWrapper(controllor.Login))

	//控制台路由组
	admin := router.Group("/admin")
	admin.GET("/index", controllor.Index)
	admin.Use(middleware.Check())
	{
		admin.GET("/info", e.ErrorWrapper(controllor.AdminInfo))

		//api := admin.Group("/api")
		//{
		//	api.GET("/index", e.ErrorWrapper(controllor.xxxx))
		//}
	}

	return router
}
