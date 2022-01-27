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
	router.GET("/isregistered", isFirstLogin)
	router.POST("/add-student-id", addStudentId)
	router.POST("/addinfo", addUserBaseInfo)
	router.GET("/get-tags", GetAllTags)
}
func initFlagApi() {
	router.POST("/create-flag", userCreateFlag)
	router.GET("/doing-flag", doingFlag)
	router.GET("/get-doing-flag", getUserDoingFlagger)
	router.GET("/get-finished-flag", getUserFinishedFlagger)
}

func Run() {
	router.Run("127.0.0.1:8080")
}
