package dbhandlers

import "hash-roar/flagger-backend/models"

func getUserDoingFlagger(uid int) ([]models.DoingFlaggersQuery, error) {
	var queryData []models.DoingFlaggersQuery
	result := db.Model(&models.UserFlagger{}).
		Select("user_flaggers.flag_sum,flaggers.id,flaggers.end_time").
		Joins("left join flaggers on user_flaggers.fid = flaggers.id").
		Where("user_flaggers.uid = ?", uid).
		Find(&queryData)
	if result.Error != nil {
		return nil, result.Error
	}
	return queryData, nil
}
