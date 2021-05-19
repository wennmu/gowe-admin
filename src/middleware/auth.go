package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wennmu/gowe-admin.git/src/config"
	"strconv"
	"strings"
)

//校验登录状态
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		tokenSlice := strings.Split(bearToken, " ")
		token, err := jwt.Parse(tokenSlice[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(config.APP_SECRET), nil
		})
		if err != nil {
			panic(errors.New("Illegal token"))
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		var uid int64
		if ok && token.Valid {
			uid, err = strconv.ParseInt(fmt.Sprintf("%.f", claims["uid"]), 10, 64)
			if err != nil || uid <= 0 {
				panic(errors.New("Illegal user"))
			}
		}
		c.Set("uid", uid)

		c.Next()
	}
}
