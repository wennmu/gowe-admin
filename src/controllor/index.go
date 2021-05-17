package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//控制台首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}
