package models

import "time"

type DoingFlaggersQuery struct {
	FlagSum      int
	Id           int
	LastFlagTime time.Time
	EndTime      time.Time
}
