package routers

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	initLoginApi()
	initFlagApi()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
}

func initLoginApi() {
	router.POST("/login", Login)
	router.POST("/addinfo", addUserBaseInfo)
}
func initFlagApi() {
	router.POST("/create-flag", userCreateFlag)
	router.GET("/doing-flag", doingFlag)
}

func Run() {
	router.Run("127.0.0.1:8080")
}
