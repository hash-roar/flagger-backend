package dbhandlers

import (
	"flagger-backend/appconfig"
	"flagger-backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// TODO: migrate to mysql
func init() {
	dsn := appconfig.AppConfig.Dsn
	dbTemp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = dbTemp
	migrate()
}

func migrate() {
	db.AutoMigrate(&models.UserBaseInfo{})
	db.AutoMigrate(&models.UserFlaggerInfo{})
	db.AutoMigrate(&models.UserFlagger{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.UserIntreTag{})
	db.AutoMigrate(&models.UserSocialTrend{})
	db.AutoMigrate(&models.Flagger{})
	db.AutoMigrate(&models.FlaggerTag{})
}
