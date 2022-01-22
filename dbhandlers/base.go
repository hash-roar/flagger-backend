package dbhandlers

import (
	"hash-roar/flagger-backend/appconfig"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := appconfig.AppConfig.Dsn
	dbTemp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = dbTemp
}
