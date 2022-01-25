package dbhandlers

import (
	"flagger-backend/models"
	"time"
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
	return data.Tid, result.Error
}

func AddUserFlagger(uid int, fid int) (int, error) {
	userFlagger := &models.UserFlagger{Uid: uid, Fid: fid, FlagSum: 1, LastFlagTime: time.Now(), Status: 1, SequentialFlagTimes: 1}
	result := db.Create(userFlagger)
	return userFlagger.Id, result.Error
}

func GetTagByTitle(title string) (int, error) {
	tag := &models.Tag{}
	result := db.Where("title = ?", title).First(tag)
	return tag.Tid, result.Error
}

func getDoingFlagger(uid int, fid int) (*models.UserFlagger, error) {
	queryData := &models.UserFlagger{}
	if err := db.Where("uid = ? AND fid = ?", uid, fid).
		First(queryData).Error; err != nil {
		return nil, err
	}
	return queryData, nil
}

func AddFlaggerTagInfo(data *models.FlaggerTag) error {
	return db.Create(data).Error
}
