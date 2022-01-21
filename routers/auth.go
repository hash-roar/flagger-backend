package routers

import (
	"hash-roar/flagger-backend/appconfig"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claim struct {
	Openid string
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	code := c.PostForm("code")
	openid := c.Request.Header.Get("X-WX-OPENID")

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
