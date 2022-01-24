package dbhandlers

import (
	"flagger-backend/models"
)

func AddUserBaseInfo(uid int, sex int, grade int, major int) (int, error) {
	result := db.Table("user_base_infos").
		Where("uid = ?", uid).
		Updates(map[string]interface{}{"sex": sex, "grade": grade, "major": major})
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

func getDoingFlagger(uid int, fid int) (*models.UserFlagger, error) {
	queryData := &models.UserFlagger{}
	if err := db.Where("uid = ? AND fid = ?", uid, fid).
		First(queryData).Error; err != nil {
		return nil, err
	}
	return queryData, nil
}
