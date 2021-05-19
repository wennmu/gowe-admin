package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wennmu/gowe-admin.git/src/router"
	"io"
	"os"
)

func main() {
	dir, _ := os.Getwd()

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create(dir + "/src/log/gowe.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.SetMode("debug")

	r := router.Init()

	r.Run(":28080")
}
