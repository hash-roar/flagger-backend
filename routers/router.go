package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

func init() {
	initLogApi()
}


func initLogApi() {
	router.GET("/login", Login)
}

func Run() {
	router.Run("127.0.01:8000")
}
