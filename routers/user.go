package routers

import (
	"hash-roar/flagger-backend/dbhandlers"
	"hash-roar/flagger-backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserBaseInfo(c *gin.Context) {
	formData := &models.FirstLoginInfo{}
	if err := c.ShouldBindJSON(formData); err != nil {
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
	userBaseInfo := &models.UserBaseInfo{}
	userSocialtrend := &models.UserSocialTrend{}
	userBaseInfo.Grade = formData.Grade
	userBaseInfo.Major = formData.Major
	userBaseInfo.Sex = formData.Sex
	userSocialtrend.SocialTrend = formData.Socialtendency
	userSocialtrend.EnvTrend = formData.Environment
	if _, err := dbhandlers.AddUserBaseInfo(userBaseInfo); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurs",
		})
		return
	}
	if _, err := dbhandlers.AddUserSocailTrend(userSocialtrend); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurs",
		})
		return
	}
	if formData.Interestedtag != nil {
		if err := dbhandlers.AddUserIntreTags(uid, formData.Interestedtag); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurs",
			})
			return
		}
	}
	if formData.CreatedTag != "" {
		tagTemp := models.Tag{TiTle: formData.CreatedTag, CreatorId: uid}
		_, err := dbhandlers.AddTag(&tagTemp)
		userIntreTagTemp := models.UserIntreTag{Uid: uid, TagTitle: formData.CreatedTag}
		_, err = dbhandlers.AddUserIntreTag(&userIntreTagTemp)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurs",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "add user info ok",
	})

}
