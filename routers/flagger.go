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
	var returnData []models.UserDoingFlagger
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	queryData, err := dbhandlers.GetUserDoingFlagger(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}

	for _, v := range queryData {
		userFlagedAvatarUrls, userFlagedNum, err := dbhandlers.
			GetDoingFlaggerUserInfo(v.Id)
		if err != nil {
			log.Println(err)
		}
		temp := models.UserDoingFlagger{}
		temp.FinishedAvatarUrl = userFlagedAvatarUrls
		temp.FinishedNum = userFlagedNum
		temp.FlaggerTitle = v.Title
		temp.FlaggerProgress = strconv.Itoa(v.FlagSum) + "/" + strconv.Itoa(v.ShouldFlagSum)
		temp.Fid = v.Id
		returnData = append(returnData, temp)
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, returnData)
}
func getUserFinishedFlagger(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	var returnData []models.UserDoingFlagger
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	queryData, err := dbhandlers.GetUserFinishedFlagger(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}

	for _, v := range queryData {
		userFlagedAvatarUrls, userFlagedNum, err := dbhandlers.
			GetFinishedFlaggerUserInfo(v.Id)
		if err != nil {
			log.Println(err)
		}
		temp := models.UserDoingFlagger{}
		temp.Fid = v.Id
		temp.FinishedAvatarUrl = userFlagedAvatarUrls
		temp.FinishedNum = userFlagedNum
		temp.FlaggerTitle = v.Title
		temp.FlaggerProgress = strconv.Itoa(v.FlagSum) + "/" + strconv.Itoa(v.ShouldFlagSum)
		returnData = append(returnData, temp)
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, returnData)
}

func doingFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	fid := int(getJsonParam(c, "fid").(float64))
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = dbhandlers.DoingFlag(uid, fid); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "打卡成功",
	})

}

func abandonFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	fid := int(getJsonParam(c, "fid").(float64))
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	if err = dbhandlers.AbandonFlag(uid, fid); err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func userCreateFlag(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	formData := &models.FormUserCreateFlag{}
	if err := c.ShouldBind(formData); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	formData.CreatorId = uid
	fid, err := dbhandlers.UserCreateFlag(formData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	_, err = dbhandlers.AddUserFlagger(uid, fid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	if formData.CreatedTag != "" {
		tagTemp := &models.Tag{CreatorId: uid, Title: formData.CreatedTag}
		tid, err := dbhandlers.AddTag(tagTemp)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		flagTagInfo := &models.FlaggerTag{Fid: fid, Tid: tid, CreateTime: time.Now()}
		err = dbhandlers.AddFlaggerTagInfo(flagTagInfo)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
	} else {
		tagTemp := &models.Tag{CreatorId: uid, Title: formData.Tag}
		tid, err := dbhandlers.AddTag(tagTemp)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		flagTagInfo := &models.FlaggerTag{Fid: fid, Tid: tid, CreateTime: time.Now()}
		err = dbhandlers.AddFlaggerTagInfo(flagTagInfo)
		if err != nil {
			log.Println(err)
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

func joinFlagGroup(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")

	// type formStruct struct {
	// 	Fid int `json:"fid" form:"fid"`
	// }
	// formData := &formStruct{}
	// c.Bind(formData)
	fid := getJsonParam(c, "fid").(float64)

	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = dbhandlers.JoinFlagger(uid, int(fid)); err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "加入成功",
	})
}

func MoreFlagger(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	// fid, _ := strconv.Atoi(c.PostForm("fid"))
	uid, err := dbhandlers.GetUidByOpenid(openid)
	var returnData []models.FindFlagger
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	allFlaggers, err := dbhandlers.GetAllFlaggers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	for _, v := range allFlaggers {
		tempFindFlagger := &models.FindFlagger{}
		tag, err := dbhandlers.GetTagTitleByFid(v.Id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		tempFindFlagger.Fid = v.Id
		tempFindFlagger.TagTitle = tag
		tempFindFlagger.FlaggerTitle = v.Title
		tempFindFlagger.ShouldFlagSum = v.ShouldFlagSum
		tempFindFlagger.Announcement = v.Announcement
		tempFindFlagger.IsMember = dbhandlers.HasJoinedFlagger(uid, v.Id)
		flaggerMemberInfo, err1 := dbhandlers.GetFlaggerMemberInfo(v.Id)
		if err1 != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		for _, memberInfo := range flaggerMemberInfo {
			if memberInfo.Uid == v.CreatorId {
				memberInfo.IsAdmin = true
			} else {
				memberInfo.IsAdmin = false
			}
		}
		tempFindFlagger.FlaggerMember = flaggerMemberInfo
		returnData = append(returnData, *tempFindFlagger)
	}
	c.JSON(http.StatusOK, returnData)

}

func GetTags(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	returnData := models.ReturnTagsInfo{}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	returnData.UserIntreTag, returnData.AllTags, err = dbhandlers.GetTags(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, returnData)
}

func SearchFlagger(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	keyWord := getJsonParam(c, "key_word").(string)
	var returnData []models.FindFlagger
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	allFlaggers, err := dbhandlers.SearchFlagger(keyWord)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	for _, v := range allFlaggers {
		tempFindFlagger := &models.FindFlagger{}
		tag, err := dbhandlers.GetTagTitleByFid(v.Id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		tempFindFlagger.Fid = v.Id
		tempFindFlagger.TagTitle = tag
		tempFindFlagger.FlaggerTitle = v.Title
		tempFindFlagger.ShouldFlagSum = v.ShouldFlagSum
		tempFindFlagger.Announcement = v.Announcement
		tempFindFlagger.IsMember = dbhandlers.HasJoinedFlagger(uid, v.Id)
		flaggerMemberInfo, err1 := dbhandlers.GetFlaggerMemberInfo(v.Id)
		if err1 != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": "服务端错误",
			})
			return
		}
		for _, memberInfo := range flaggerMemberInfo {
			if memberInfo.Uid == v.CreatorId {
				memberInfo.IsAdmin = true
			} else {
				memberInfo.IsAdmin = false
			}
		}
		tempFindFlagger.FlaggerMember = flaggerMemberInfo
		returnData = append(returnData, *tempFindFlagger)
	}
	c.JSON(http.StatusOK, returnData)
}

func GetAllTags(c *gin.Context) {
	tags, err := dbhandlers.GetAllTags()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	tagStr := make([]string, 0)
	for _, v := range tags {
		tagStr = append(tagStr, v.Title)
	}
	c.JSON(http.StatusOK, tagStr)
}

func getJsonParam(c *gin.Context, key string) (result interface{}) {
	vMaps := make(map[string]interface{}, 0)
	c.BindJSON(&vMaps)
	return vMaps[key]
}

func getFlagInfoByFid(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	// fid, _ := strconv.Atoi(c.PostForm("fid"))
	fid := int(getJsonParam(c, "fid").(float64))
	uid, err := dbhandlers.GetUidByOpenid(openid)

	returnData := models.FlaggerInfo{}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	flagger, err := dbhandlers.GetFlaggerByFid(fid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	tempFindFlagger := &models.FlaggerInfo{}
	tag, err := dbhandlers.GetTagTitleByFid(flagger.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	tempFindFlagger.Fid = flagger.Id
	tempFindFlagger.TagTitle = tag
	tempFindFlagger.FlaggerTitle = flagger.Title
	tempFindFlagger.ShouldFlagSum = flagger.ShouldFlagSum
	tempFindFlagger.Announcement = flagger.Announcement
	tempFindFlagger.IsMember = dbhandlers.HasJoinedFlagger(uid, flagger.Id)
	flaggerMemberInfo, err1 := dbhandlers.GetFlaggerMemberInfoPlus(flagger.Id)
	if err1 != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	for _, memberInfo := range flaggerMemberInfo {
		if memberInfo.Uid == flagger.CreatorId {
			memberInfo.IsAdmin = true
		} else {
			memberInfo.IsAdmin = false
		}
	}
	tempFindFlagger.FlaggerMember = flaggerMemberInfo
	returnData = *tempFindFlagger

	c.JSON(http.StatusOK, returnData)

}
