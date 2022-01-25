package models

import "time"

type DoingFlaggersQuery struct {
	FlagSum      int
	Id           int
	Title        string
	LastFlagTime time.Time
	EndTime      time.Time
}
