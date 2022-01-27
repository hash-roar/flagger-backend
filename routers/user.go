package routers

import (
	"flagger-backend/dbhandlers"
	"flagger-backend/models"
	"flagger-backend/tools"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserBaseInfo(c *gin.Context) {
	formData := &models.FormUserBaseInfo{}
	if err := c.Bind(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	userSocialtrend := &models.UserSocialTrend{}
	userSocialtrend.Uid = uid
	socialTrendArr := [...]int{formData.Socialtendency}
	userSocialtrend.SocialTrend = tools.SwitchArrayToNum(socialTrendArr[:])
	eneTrendArr := [...]int{formData.Environment}
	userSocialtrend.EnvTrend = tools.SwitchArrayToNum(eneTrendArr[:])
	if _, err := dbhandlers.AddUserBaseInfo(uid, formData.Sex, formData.Grade, formData.Major); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "服务端错误",
		})
		return
	}
	if _, err := dbhandlers.AddUserSocailTrend(userSocialtrend); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "服务端错误",
		})
		return
	}
	if formData.InterestedTag != nil {
		if err := dbhandlers.AddUserIntreTags(uid, formData.InterestedTag); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "服务端错误",
			})
			return
		}
	}
	if formData.CreatedTag != "" {
		tagTemp := models.Tag{Title: formData.CreatedTag, CreatorId: uid}
		tid, err := dbhandlers.AddTag(&tagTemp)
		userIntreTagTemp := models.UserIntreTag{Uid: uid, TagTitle: formData.CreatedTag, Tid: tid}
		_, err = dbhandlers.AddUserIntreTag(&userIntreTagTemp)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "服务端错误",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加用户信息成功",
	})
}

func UserInfo(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	returnData := &models.UserInfo{}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	err = dbhandlers.GetUserBaseInfo(uid, returnData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	returnData.CredenceValue, err = dbhandlers.GetUserCredenceValue(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	returnData.HaveFlaged, err = dbhandlers.GetUserHaveFlaggedSun(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	returnData.ShouldFlagSum, err = dbhandlers.GetUserShouldFlaggedSum(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, returnData)
}

func isFirstLogin(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	if dbhandlers.IsRegistered(openid) {
		c.JSON(http.StatusOK, gin.H{
			"is_registered": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"is_registered": false,
		})
	}
}
