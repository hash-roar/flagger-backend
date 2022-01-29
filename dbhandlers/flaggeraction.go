package dbhandlers

import (
	"errors"
	"flagger-backend/models"
	"flagger-backend/tools"
	"time"

	"gorm.io/gorm"
)

func DoingFlag(uid int, fid int) error {
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
	} else {
		userFlagger.FlagSum += 1
		userFlagger.SequentialFlagTimes = 1
	}
	userFlagger.LastFlagTime = time.Now()
	if err = AddFlaggerTotalSum(fid); err != nil {
		return err
	}
	if err = AddUserReputation(uid, userFlagger.SequentialFlagTimes); err != nil {
		return err
	}
	return db.Save(userFlagger).Error
}

func GetUserDoingFlagger(uid int) ([]models.DoingFlaggersQuery, error) {
	var queryData []models.DoingFlaggersQuery
	result := db.Model(&models.UserFlagger{}).
		Select("user_flaggers.flag_sum,user_flaggers.last_flag_time,flaggers.id,flaggers.should_flag_sum,flaggers.title").
		Joins("left join flaggers on user_flaggers.fid = flaggers.id").
		Where("user_flaggers.uid = ?", uid).
		Where("user_flaggers.last_flag_time < ?", tools.GetTodayStartTime()).
		Where("user_flaggers.status = ?", 1).
		Find(&queryData)
	if result.Error != nil {
		return nil, result.Error
	}
	return queryData, nil
}

func GetFinishedFlaggerUserInfo(fid int) (flagedAvatarUrl []string, hadFlagedNum int, err error) {
	type queryStruct struct {
		AvatarUrl string
	}
	var queryData []queryStruct
	err = db.Table("user_flaggers").
		Select("user_base_infos.avatar_url").
		Joins("left join user_base_infos on user_flaggers.uid = user_base_infos.uid").
		Where("user_flaggers.fid = ?", fid).
		Where("user_flaggers.status = ?", 1).
		Where("user_flaggers.last_flag_time  BETWEEN ? AND ?", tools.GetTodayStartTime(), time.Now()).
		Find(&queryData).Error
	hadFlagedNum = len(queryData)
	for _, v := range queryData {
		flagedAvatarUrl = append(flagedAvatarUrl, v.AvatarUrl)
	}
	return
}

func GetDoingFlaggerUserInfo(fid int) (flagedAvatarUrl []string, hadFlagedNum int, err error) {
	type queryStruct struct {
		AvatarUrl string
	}
	var queryData []queryStruct
	err = db.Table("user_flaggers").
		Select("user_base_infos.avatar_url").
		Joins("left join user_base_infos on user_flaggers.uid = user_base_infos.uid").
		Where("user_flaggers.fid = ?", fid).
		Where("user_flaggers.last_flag_time  < ?", tools.GetTodayStartTime()).
		Find(&queryData).Error
	hadFlagedNum = len(queryData)
	for _, v := range queryData {
		flagedAvatarUrl = append(flagedAvatarUrl, v.AvatarUrl)
	}
	return
}

func UserCreateFlag(data *models.FormUserCreateFlag) (int, error) {
	flaggerTemp := &models.Flagger{}
	flaggerTemp.Announcement = data.Announcement
	flaggerTemp.JoinAuth = tools.GetAuthNum(data.JoinAuth)
	flaggerTemp.Frequency = data.Frequency
	flaggerTemp.MaxGroupMember = data.MaxGroupMember
	flaggerTemp.Title = data.Title
	flaggerTemp.EndTime = time.Now().AddDate(0, 0, data.EndTime)
	flaggerTemp.ShouldFlagSum = data.EndTime
	flaggerTemp.CreatorId = data.CreatorId

	result := db.Table("flaggers").Create(flaggerTemp)
	return flaggerTemp.Id, result.Error
}

func JoinFlagger(uid int, fid int) error {
	// type queryStruct struct {
	// 	JoinAuth uint64
	// }
	type queryStruct2 struct {
		Grade int
	}
	// tempQueryStruct := &queryStruct{}
	if db.Where("uid = ? AND fid = ?", uid, fid).
		Find(&models.UserFlagger{}).RowsAffected != 0 {
		return nil
	}
	flagger := &models.Flagger{}
	tempQueryStruct2 := &queryStruct2{}
	err := db.Table("flaggers").
		Where("id = ?", fid).
		First(flagger).Error
	if err != nil {
		return err
	}
	err = db.Table("user_base_infos").Select("grade").First(tempQueryStruct2).Error
	if err != nil {
		return err
	}
	if !tools.IsAuthorized(tempQueryStruct2.Grade, flagger.JoinAuth) {
		return errors.New("没有权限")
	}
	userFlagger := &models.
		UserFlagger{Uid: uid, Fid: fid, FlagSum: 1, SequentialFlagTimes: 1, LastFlagTime: time.Now(), Status: 1}
	err = db.Create(userFlagger).Error
	if err != nil {
		return err
	}
	err = AddFlaggerTotalSum(fid)
	return err
}

func AbandonFlag(uid int, fid int) error {
	type queryStruct struct {
		CreatorId int
	}
	queryData := &queryStruct{}
	err := db.Table("flaggers").
		Where("id = ?", fid).
		Select("creator_id").First(queryData).Error
	if err != nil {
		return err
	}
	if uid == queryData.CreatorId {
		if err = db.Table("flaggers").
			Where("id = ?", fid).
			Delete(&models.Flagger{}).Error; err != nil {
			return err
		}
		if err = db.Table("user_flaggers").
			Where("fid = ?", fid).
			Delete(&models.UserFlagger{}).Error; err != nil {
			return nil
		}
	} else {
		if err = db.Table("user_flaggers").
			Where("fid = ? AND uid = ?", fid, uid).Delete(&models.UserFlagger{}).Error; err != nil {
			return err
		}
	}
	if err = SubUserCredence(uid); err != nil {
		return err
	}
	return nil
}

func HasJoinedFlagger(uid int, fid int) bool {
	return db.Where("uid = ? AND fid = ?", uid, fid).Find(&models.UserFlagger{}).RowsAffected != 0
}

func SearchFlagger(keyWord string) (result []models.Flagger, err error) {
	err = db.Where("title LIKE ?", "%"+keyWord+"%").Find(&result).Error
	return
}

func SubUserCredence(uid int) error {
	return db.Table("user_flagger_infos").
		Where("uid = ?", uid).
		Update("credence_value", gorm.Expr("credence_value - ?", 1)).Error
}

func GetUserHistory(uid int) (result []models.Flagger, err error) {
	err = db.Model(&models.Flagger{}).
		Joins("left join user_flaggers on user_flaggers.fid = flaggers.id").
		Where("user_flaggers.uid = ?", uid).
		Find(&result).Error
	return
}
