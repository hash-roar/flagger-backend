package routers

import (
	"flagger-backend/dbhandlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserDoingFlagger(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(uid)
}

func doingFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	openid := c.PostForm("header")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err!=dbhandlers.doingFlag(uid,fid) {
		
	}

}
