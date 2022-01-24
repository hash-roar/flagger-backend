package dbhandlers

import (
	"flagger-backend/models"
	"flagger-backend/tools"
	"time"
)

func doingFlag(uid int, fid int) error {
	userFlagger, err := getDoingFlagger(uid, fid)
	if err != nil {
		return err
	}
	if tools.IsToday(userFlagger.LastFlagTime) {
		return nil
	}
	if tools.IsYesterday(userFlagger.LastFlagTime) {
		userFlagger.FlagSum += 1
		userFlagger.SequentialFlagTimes += 1
	}
	userFlagger.FlagSum += 1
	userFlagger.SequentialFlagTimes = 1
	userFlagger.LastFlagTime = time.Now()
	return db.Save(userFlagger).Error
}

func getUserDoingFlagger(uid int) ([]models.DoingFlaggersQuery, error) {
	var queryData []models.DoingFlaggersQuery
	result := db.Model(&models.UserFlagger{}).
		Select("user_flaggers.flag_sum,user_flaggers.last_flag_time,flaggers.id,flaggers.end_time").
		Joins("left join flaggers on user_flaggers.fid = flaggers.id").
		Where("user_flaggers.uid = ?", uid).
		Where("user_flaggers.last_flag_time ").
		Find(&queryData)
	if result.Error != nil {
		return nil, result.Error
	}
	return queryData, nil
}
