package routers

import (
	"flagger-backend/appconfig"
	"flagger-backend/dbhandlers"
	"flagger-backend/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claim struct {
	Openid string
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	formData := &models.FormLoginInfo{}
	if err := c.Bind(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	openid := c.Request.Header.Get("X-WX-OPENID")
	formData.Openid = openid
	if err := dbhandlers.AddUserLoginInfo(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tokenStr, err := generateToken(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   tokenStr,
		"uid":     uid,
		"message": "登录成功",
	})
}

func addStudentId(c *gin.Context) {
	type formStruct struct {
		StudentId string `json:"student_id"`
		Password  string `json:"password"`
		AvatarUrl string `json:"avatar_url"`
		Nickname  string `json:"nickname"`
	}
	openid := c.Request.Header.Get("X-WX-OPENID")
	formData := &formStruct{}
	c.Bind(formData)
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	err = dbhandlers.AddStudentId(uid, formData.StudentId, formData.Password,
		formData.AvatarUrl, formData.Nickname)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

func getToken(c *gin.Context) {
	openid := c.Request.Header.Get("X-WX-OPENID")
	uid, err := dbhandlers.GetUidByOpenid(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	tokenStr, err := generateToken(openid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "服务端错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
		"uid":   uid,
	})

}

func generateToken(id string) (string, error) {
	// maxage := appconfig.AppConfig.MaxAge
	claim := &Claim{
		Openid: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(appconfig.AppConfig.JwtSec))
}

func authToken(tokenStr string) (openid string, ok bool) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(appconfig.AppConfig.JwtSec), nil
	})
	if err != nil || !token.Valid {
		log.Println(err)
		return "", false
	}
	// if claim.ExpiresAt {

	// }
	return claim.Openid, true
}
