package controllor

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wennmu/gowe-admin.git/pkg/e"
	"github.com/wennmu/gowe-admin.git/src/common"
	"github.com/wennmu/gowe-admin.git/src/config"
	"time"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) (interface{}, error) {
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, e.AppError{Code: common.ERROR, Msg: err.Error()}
	}

	//TODO 获取用户ID和用户ACCESS_SECRET
	uid := 1

	atCliams := jwt.MapClaims{}

	atCliams["authorized"] = true
	atCliams["userId"] = uid
	atCliams["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atCliams)
	token, err := at.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return nil, e.AppError{Code: common.ERROR, Msg: err.Error()}
	}

	return map[string]string{
		"token": token,
	}, nil
}
