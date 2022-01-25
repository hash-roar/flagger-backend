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
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "登录成功",
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
	return token.SignedString(appconfig.AppConfig.JwtSec)
}

func authToken(tokenStr string) (openid string, ok bool) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return appconfig.AppConfig.JwtSec, nil
	})
	if err != nil || !token.Valid {
		log.Println(err)
		return "", false
	}
	return claim.Openid, true
}
