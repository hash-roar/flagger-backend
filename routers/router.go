package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	noAuthApi()
	// router.Use(AuthMidware())
	initLoginApi()
	initFlagApi()
	initUserApi()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
}

func AuthMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authentication")
		openidFromHeader := c.Request.Header.Get("X-WX-OPENID")
		if openid, ok := authToken(token); ok == true && openid == openidFromHeader {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "认证失败"})
			c.Abort()
		}
	}
}

func initLoginApi() {
	router.POST("/add-student-id", addStudentId)
	router.POST("/addinfo", addUserBaseInfo)
}
func initFlagApi() {
	router.GET("/get-tags", GetAllTags)
	router.POST("/create-flag", userCreateFlag)
	router.POST("/doing-flag", doingFlag)
	router.GET("/get-doing-flag", getUserDoingFlagger)
	router.GET("/get-finished-flag", getUserFinishedFlagger)
	router.GET("/join-flag", joinFlagGroup)
	router.POST("/abandon-flag", abandonFlag)
	router.GET("/get-intre_tags", GetTags)
}
func initUserApi() {
	router.GET("/get-user-info", UserInfo)
	router.GET("/get-user-history", GetUserHistory)
	router.POST("/save-user-info", SaveUserInfo)
}

func noAuthApi() {
	router.GET("/get-flags", MoreFlagger)
	router.POST("/login", Login)
	router.GET("/isregistered", isFirstLogin)
	router.POST("/search-flag", SearchFlagger)
	router.GET("/get-token", getToken)
	router.POST("/flaginfo", getFlagInfoByFid)
}

// func authApi()  {

// }

func Run() {
	router.Run(":80")
}
