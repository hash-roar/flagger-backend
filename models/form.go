package models

type FirstLoginInfo struct {
	Sex            int `form:"sex"`
	Grade          int
	Major          int
	Interestedtag  []string
	CreatedTag     string
	Environment    int
	Socialtendency int
}
