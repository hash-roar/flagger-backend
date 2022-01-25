package routers

import (
	"flagger-backend/dbhandlers"
	"flagger-backend/models"
	"log"
	"net/http"
	"strconv"
	"time"

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
	queryData, err := dbhandlers.GetUserDoingFlagger(uid)
	if err != nil {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
	}
	userData,err:= dbhandlers.
}

func doingFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	fid, _ := strconv.Atoi(c.PostForm("fid"))
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != dbhandlers.DoingFlag(uid, fid) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "打卡成功",
	})

}

func userCreateFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	formData := &models.FormUserCreateFlag{}
	if err := c.ShouldBind(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	fid, err := dbhandlers.UserCreateFlag(formData)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	_, err = dbhandlers.AddUserFlagger(uid, fid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	if formData.CreatedTag != "" {
		tagTemp := &models.Tag{CreatorId: uid, Title: formData.CreatedTag}
		tid, err := dbhandlers.AddTag(tagTemp)
		flagTagInfo := &models.FlaggerTag{Fid: fid, Tid: tid, CreateTime: time.Now()}
		err = dbhandlers.AddFlaggerTagInfo(flagTagInfo)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
	} else {
		tid, err := dbhandlers.GetTagByTitle(formData.Tag)
		flagTagInfo := &models.FlaggerTag{Fid: fid, Tid: tid, CreateTime: time.Now()}
		err = dbhandlers.AddFlaggerTagInfo(flagTagInfo)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}
