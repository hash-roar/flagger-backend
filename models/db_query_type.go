package models

import "time"

//TODO: combine tuple with crud operation

type DoingFlaggersQuery struct {
	Id            int
	FlagSum       int
	Title         string
	LastFlagTime  time.Time
	ShouldFlagSum int
}
