package models

import "time"

type DoingFlaggersQuery struct {
	Id            int
	FlagSum       int
	Title         string
	LastFlagTime  time.Time
	ShouldFlagSum int
}

