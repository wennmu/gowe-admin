package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//校验登录状态
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		//鉴权
		fmt.Println("鉴权...")
		c.Next()
	}
}
