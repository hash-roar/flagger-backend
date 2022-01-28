package dbhandlers

import (
	"errors"
	"flagger-backend/models"
	"flagger-backend/tools"
	"time"
)

func GetUidByOpenid(openid string) (int, error) {
	userInfo := &models.UserBaseInfo{}
	result := db.Where("openid = ?", openid).First(userInfo)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("no such uer")
	}
	return userInfo.Uid, nil
}

func AddUserLoginInfo(data *models.FormLoginInfo) error {
	userBaseInfo := &models.UserBaseInfo{}
	userBaseInfo.AvatarUrl = data.AvatarUrl
	userBaseInfo.Nickname = data.Nickname
	// userBaseInfo.StudentId = data.StudentId
	// userBaseInfo.Password = data.Password
	userBaseInfo.Openid = data.Openid
	if err := db.Table("user_base_infos").Create(userBaseInfo).Error; err != nil {
		return err
	}
	if err := db.Create(&models.UserFlaggerInfo{
		Uid: userBaseInfo.Uid, CredenceValue: 100, ReputationValue: 0}).
		Error; err != nil {
		return err
	}
	return nil
}

func AddUserIntreTags(uid int, tags []string) error {
	for _, v := range tags {
		userIntreTagTemp := &models.UserIntreTag{Uid: uid, TagTitle: v, CreateTime: time.Now()}
		if _, err := AddUserIntreTag(userIntreTagTemp); err != nil {
			return err
		}
	}
	return nil
}

func AddStudentId(uid int, studentId string, password string) error {
	return db.Table("user_base_infos").
		Where("uid = ?", uid).
		Updates(map[string]interface{}{"student_id": studentId, "password": password}).
		Error
}

func GetFlaggerMemberInfo(fid int) ([]models.FlaggerGroupMemberInfo, error) {
	type queryStruct struct {
		Uid       int
		AvatarUrl string
		Nickname  string
		FlagSum   int
	}
	var queryData []queryStruct
	var flaggerMemberInfo []models.FlaggerGroupMemberInfo
	err := db.Table("user_flaggers").
		Joins("left join user_base_infos on user_flaggers.uid = user_base_infos.uid").
		Where("user_flaggers.fid = ?", fid).
		Select("user_base_infos.avatar_url", "user_base_infos.nickname", "user_flaggers.flag_sum", "user_base_infos.uid").
		Find(&queryData).Error
	if err != nil {
		return nil, err
	}
	for _, v := range queryData {
		var userIntreFlags []models.UserIntreTag
		var userIntreFlagsString []string
		err = db.Where("uid = ?", v.Uid).Find(&userIntreFlags).Error
		if err != nil {
			return nil, err
		}
		for _, v := range userIntreFlags {
			userIntreFlagsString = append(userIntreFlagsString, v.TagTitle)
		}
		flaggerMemberInfo = append(flaggerMemberInfo, models.FlaggerGroupMemberInfo{
			AvatarUrl:    v.AvatarUrl,
			Nickname:     v.Nickname,
			FlagSum:      v.FlagSum,
			UserIntreTag: userIntreFlagsString,
			Uid:          v.Uid,
		})
	}
	return flaggerMemberInfo, nil
}

func GetTags(uid int) (UserIntreTag []string, AllTags []string, err error) {
	var userIntreTags []models.UserIntreTag
	var allTags []models.Tag
	if err = db.Where("uid = ?", uid).Find(&userIntreTags).Error; err != nil {
		return
	}
	for _, v := range userIntreTags {
		UserIntreTag = append(UserIntreTag, v.TagTitle)
	}
	if err = db.Find(&allTags).Error; err != nil {
		return
	}
	for _, v := range allTags {
		AllTags = append(AllTags, v.Title)
	}
	return
}

func GetUserBaseInfo(uid int, data *models.UserInfo) error {
	userBaseInfo := &models.UserBaseInfo{}
	err := db.Where("uid = ?", uid).
		First(userBaseInfo).Error
	data.AvatarUrl = userBaseInfo.AvatarUrl
	data.Nickname = userBaseInfo.Nickname
	data.Grade = userBaseInfo.Grade
	data.Major = userBaseInfo.Major
	return err
}

func GetUserSocialTrend(uid int, data *models.UserInfo) error {
	userSocialTrend := &models.UserSocialTrend{}
	err := db.Where("uid = ?", uid).First(userSocialTrend).Error
	data.UserSocialTrend = tools.SwitchNumToArray(userSocialTrend.SocialTrend)
	data.Environment = tools.SwitchNumToArray(userSocialTrend.EnvTrend)
	return err
}

func GetUserCredenceValue(uid int) (int, error) {
	userFlaggersInfo := &models.UserFlaggerInfo{}
	err := db.Where("uid = ?", uid).First(userFlaggersInfo).Error
	return userFlaggersInfo.CredenceValue, err
}

func GetUserHaveFlaggedSun(uid int) (int, error) {
	var userFlaggers []models.UserFlagger
	err := db.Table("user_flaggers").
		Where("uid = ?", uid).
		Where("status = ?", 1).
		Where("last_flag_time BETWEEN ? AND ?", tools.GetTodayStartTime(), time.Now()).
		Find(&userFlaggers).Error
	return len(userFlaggers), err
}

func GetUserShouldFlaggedSum(uid int) (int, error) {
	var userFlaggers []models.UserFlagger
	err := db.Table("user_flaggers").
		Where("uid = ?", uid).
		Where("status = ?", 1).
		Find(&userFlaggers).Error
	return len(userFlaggers), err
}

func IsRegistered(openid string) bool {
	user := &models.UserBaseInfo{}
	result := db.Where("openid = ?", openid).Find(user)
	if result.RowsAffected == 0 {
		return false
	}
	return true

}

func GetFlaggerUserNum(fid int) (int, error) {
	var userFlaggers []models.UserFlagger
	result := db.Where("fid = ?", fid).Find(&userFlaggers)
	return int(result.RowsAffected), result.Error
}
