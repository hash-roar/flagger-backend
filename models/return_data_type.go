package models

type UserDoingFlagger struct {
	Fid               int      `json:"fid"`
	FlaggerTitle      string   `json:"flagger_title"`
	FinishedAvatarUrl []string `json:"finished_avatar_url"`
	FinishedNum       int      `json:"finished_num"`
	FlaggerProgress   string   `json:"flagger_progress"`
}

type FlaggerGroupMemberInfo struct {
	Uid          int
	AvatarUrl    string
	Nickname     string
	IsAdmin      bool
	FlagSum      int
	UserIntreTag []string
}

type FindFlagger struct {
	Fid           int
	FlaggerTitle  string
	TagTitle      string
	Announcement  string
	IsMember      bool
	ShouldFlagSum int
	FlaggerMember []FlaggerGroupMemberInfo
}

type FlaggerGroupMemberInfoPlus struct {
	Uid                int      `json:"uid"`
	AvatarUrl          string   `json:"avatar_url"`
	Nickname           string   `json:"nickname"`
	IsAdmin            bool     `json:"is_admin"`
	FlagSum            int      `json:"flag_sum"`
	UserIntreTag       []string `json:"user_intre_tag"`
	SequentialFlagTime int      `json:"sequential_flag_time"`
}

type FlaggerInfo struct {
	Fid           int
	FlaggerTitle  string
	TagTitle      string
	Announcement  string
	IsMember      bool
	ShouldFlagSum int
	FlaggerMember []FlaggerGroupMemberInfoPlus
}

type ReturnTagsInfo struct {
	UserIntreTag []string
	AllTags      []string
}

type UserInfo struct {
	Nickname        string `json:"nickname"`
	AvatarUrl       string `json:"avatar_url"`
	Grade           int    `json:"grade"`
	Major           int    `json:"major"`
	UserSocialTrend []int  `json:"user_social_trend"`
	Environment     []int  `json:"environment"`
	HaveFlaged      int    `json:"have_flaged"`
	ShouldFlagSum   int    `json:"should_flag_sum"`
	CredenceValue   int    `json:"credence_value"`
	// UserHistory     []FindFlagger
}
