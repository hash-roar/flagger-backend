package models

type FormUserBaseInfo struct {
	Sex            int      `form:"sex" json:"sex"`
	Grade          int      `form:"grade" json:"grade"`
	Major          int      `form:"major" json:"major"`
	InterestedTag  []string `form:"interested_tag" json:"interested_tag"`
	CreatedTag     string   `form:"created_tag" json:"created_tag"`
	Environment    int      `form:"environment" json:"environment"`
	Socialtendency int      `form:"socialtendency" json:"socialtendency"`
}

type FormLoginInfo struct {
	Openid    string `json:"openid,omitempty" form:"openid"`
	StudentId string `json:"student_id,omitempty" form:"student_id"`
	Password  string `json:"password,omitempty" form:"password"`
	Code      string `json:"code,omitempty" form:"code"`
	AvatarUrl string `json:"avatar_url,omitempty" form:"avatar_url"`
	Nickname  string `json:"nickname,omitempty" form:"nickname"`
}
