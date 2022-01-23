package routers

import (
	"hash-roar/flagger-backend/dbhandlers"
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
	

}
