package dbhandlers

import (
	"flagger-backend/models"
	"flagger-backend/tools"

	"gorm.io/gorm"
)

func AddUserBaseInfo(uid int, sex int, grade int, major int) (int, error) {
	result := db.Table("user_base_infos").
		Where("uid = ?", uid).
		Updates(map[string]interface{}{"sex": sex, "grade": grade, "major": major})
	return int(result.RowsAffected), result.Error
}

func SaveUserBaseInfo(uid int, data *models.FormSaveUserInfo) error {
	userInfo := &models.UserBaseInfo{}
	if err := db.Where("uid = ?", uid).First(userInfo).Error; err != nil {
		return err
	}
	userInfo.AvatarUrl = data.AvatarUrl
	userInfo.Nickname = data.Nickname
	userInfo.Major = data.Major
	userInfo.Grade = data.Grade
	if err := db.Model(userInfo).Where("uid = ?", uid).Updates(*userInfo).Error; err != nil {
		return err
	}
	userSocialTrend := &models.UserSocialTrend{}
	if err := db.Where("uid = ?", uid).First(userSocialTrend).Error; err != nil {
		return err
	}
	userSocialTrend.EnvTrend = tools.SwitchArrayToNum(data.Environment)
	userSocialTrend.SocialTrend = tools.SwitchArrayToNum(data.Socialtendency)
	if err := db.Model(userSocialTrend).Where("uid = ?", uid).Updates(*userSocialTrend).Error; err != nil {
		return err
	}
	return nil
}

func AddUserSocailTrend(data *models.UserSocialTrend) (int, error) {
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}

func AddUserIntreTag(data *models.UserIntreTag) (int, error) {
	tag := &models.Tag{}
	if db.Where("title = ?", data.TagTitle).First(tag).RowsAffected == 0 {
		tag.Title = data.TagTitle
		tag.CreatorId = data.Uid
		if err := db.Create(tag).Error; err != nil {
			return 0, err
		}
	}
	data.Tid = tag.Tid
	result := db.Create(data)
	return int(result.RowsAffected), result.Error
}

func AddUserFlagger(uid int, fid int) (int, error) {
	userFlagger := &models.UserFlagger{Uid: uid, Fid: fid,
		FlagSum: 0, Status: 1, SequentialFlagTimes: 0, LastFlagTime: tools.GetYesterdayStartTime()}
	result := db.Create(userFlagger)
	return userFlagger.Id, result.Error
}

func AddFlaggerTotalSum(fid int) error {
	return db.Table("flaggers").
		Where("id = ?", fid).
		Update("total_flags", gorm.Expr("total_flags + ?", 1)).Error
}

func GetTagByTitle(title string) (int, error) {
	tag := &models.Tag{}
	result := db.Where("title = ?", title).First(tag)
	return tag.Tid, result.Error
}

func GetAllFlaggers() ([]models.Flagger, error) {
	var queryData []models.Flagger
	if err := db.Find(&queryData).Error; err != nil {
		return nil, err
	}
	return queryData, nil

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

func AddTag(data *models.Tag) (int, error) {
	tag := &models.Tag{}
	if db.Where("title = ?", data.Title).First(tag).RowsAffected == 0 {
		result := db.Create(data)
		return data.Tid, result.Error
	}
	return tag.Tid, nil
}

func GetTagTitleByFid(fid int) (string, error) {
	type queryStruct struct {
		Title string
	}
	tempQueryData := &queryStruct{}
	err := db.Table("tags").
		Joins("left join flagger_tags  on flagger_tags.tid = tags.tid").
		Where("flagger_tags.fid = ?", fid).
		Select("tags.title").
		First(tempQueryData).Error
	if err != nil {
		return "", err
	}
	return tempQueryData.Title, nil
}

func GetAllTags() ([]models.Tag, error) {
	var allTags []models.Tag
	err := db.Limit(10).Find(&allTags).Error
	return allTags, err
}

func GetFlaggerByFid(fid int) (*models.Flagger, error) {
	flagger := &models.Flagger{}
	err := db.Where("id = ?", fid).First(flagger).Error
	return flagger, err
}
