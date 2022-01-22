package dbhandlers

import (
	"hash-roar/flagger-backend/models"
)

func AddUserBaseInfo(data *models.UserBaseInfo) (int, error) {
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}

func AddUserSocailTrend(data *models.UserSocialTrend) (int, error) {
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}

func AddUserIntreTag(data *models.UserIntreTag) (int, error) {
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}

func AddTag(data *models.Tag) (int, error) {
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}
