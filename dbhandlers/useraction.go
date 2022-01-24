package dbhandlers

import (
	"errors"
	"flagger-backend/models"
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
	userBaseInfo.StudentId = data.StudentId
	userBaseInfo.Password = data.Password
	if err := db.Table("user_base_infos").Create(userBaseInfo).Error; err != nil {
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
