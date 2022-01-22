package models

import (
	"time"
)

type UserBaseInfo struct {
	Uid   int `gorm:"primaryKey"`
	Sex   int
	Grade int
	Major int
}

type UserFlaggerInfo struct {
	Uid             int `gorm:"primaryKey"`
	CredenceValue   int
	ReputationValue int
}

type UserFlagger struct {
	Id                  int `gorm:"primaryKey"`
	Uid                 int
	Fid                 int
	FlagSum             int
	LastFlagTime        time.Time
	SequentialFlagTimes int
	staus               int
}

type UserIntreTag struct {
	Uid        int
	Tid        int
	TagTitle   string
	CreateTime time.Time
}

type UserSocialTrend struct {
	Uid         int `gorm:"primaryKey"`
	EnvTrend    int
	SocialTrend int
}

type FlaggerTag struct {
	Fid        int `gorm:"primaryKey"`
	Tid        int
	CreateTime time.Time
}

type Flagger struct {
	Id              int    `gorm:"primaryKey" json:"id,omitempty"`
	Title           string `json:"title,omitempty"`
	Type            int    `json:"type,omitempty"`
	MaxGroupMember  int    `json:"max_group_member,omitempty"`
	GroupMemberCtrl int    `json:"group_member_ctrl,omitempty"`
	Tags            string `json:"tags,omitempty"`
	Frequency       int    `json:"frequency,omitempty"`
}

type Tag struct {
	Tid         int    `json:"tid,omitempty"`
	TiTle       string `json:"ti_tle,omitempty"`
	Description string `json:"description,omitempty"`
	CreatorId   int    `json:"creator_id,omitempty"`
}
