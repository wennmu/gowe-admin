package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		//生成请求ID
		c.Set("reqid", uuid.New())

		c.Next()
	}
}
