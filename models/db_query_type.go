package models

import "time"

type DoingFlaggersQuery struct {
	FlagSum int
	Id      int
	EndTime time.Time
}
