package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBaseInfo struct {
	Uid       int    `gorm:"primarykey"`
	Openid    string `gorm:"index:idx,unique"`
	Sex       int
	Grade     int
	Major     int
	AvatarUrl string
	Nickname  string
	SecretKey string
	StudentId string
	Password  string
}

type UserFlaggerInfo struct {
	Uid             int `gorm:"primaryKey"`
	CredenceValue   int
	ReputationValue int
}

type UserFlagger struct {
	Id                  int `gorm:"primarykey"`
	Uid                 int
	Fid                 int
	FlagSum             int
	LastFlagTime        time.Time
	SequentialFlagTimes int
	Status              int // 0 give up 1 doing 2 finish
}

type UserIntreTag struct {
	Uid        int
	Tid        int
	TagTitle   string
	CreateTime time.Time
}

type UserSocialTrend struct {
	Uid         int `gorm:"primaryKey"`
	EnvTrend    uint64
	SocialTrend uint64
}

type FlaggerTag struct {
	Fid        int `gorm:"primaryKey"`
	Tid        int
	CreateTime time.Time
}

type Flagger struct {
	Id             int    `gorm:"primarykey" json:"id,omitempty"`
	Title          string `json:"title,omitempty"`
	Type           int    `json:"type,omitempty"`
	MaxGroupMember int    `json:"max_group_member,omitempty"`
	// 0 无限制 1 :3-5 ,2 : 5-10,3 : 10-20
	GroupMemberCtrl int    `json:"group_member_ctrl,omitempty"`
	Tags            string `json:"tags,omitempty"`
	Frequency       int    `json:"frequency,omitempty"`
	// 1-7
	Announcement string `json:"announcement,omitempty"`
	TotalFlags   int    `json:"total_flags,omitempty"` //总打卡次数
	FlagStatus   int    `json:"flag_status,omitempty"`
	// 0 doing 1 finish
	JoinAuth      uint64 `json:"join_auth,omitempty"`
	ShouldFlagSum int
	EndTime       time.Time `json:"end_time,omitempty"`
	CreatorId     int
}

type Tag struct {
	Tid         int    `json:"tid,omitempty" gorm:"primarykey"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	CreatorId   int    `json:"creator_id,omitempty"`
}

type UserHistory struct {
	gorm.Model
	Uid int
	Fid int
}
