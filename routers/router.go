package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

func init() {
	initLoginApi()
}

func initLoginApi() {
	router.POST("/login", Login)
}

func Run() {
	router.Run("127.0.0.1:8000")
}
