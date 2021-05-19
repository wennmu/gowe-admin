package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//控制台首页
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uid": c.Value("uid"),
	})
}
