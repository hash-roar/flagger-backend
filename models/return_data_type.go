package models

type UserDoingFlagger struct {
	Fid               int      `json:"fid"`
	FlaggerTitle      string   `json:"flagger_title"`
	FinishedAvatarUrl []string `json:"finished_avatar_url"`
	FinishedNum       int      `json:"finished_num"`
	FlaggerProgress   string   `json:"flagger_progress"`
}

type FlaggerGroupMemberInfo struct {
	Uid          int      `json:"uid"`
	AvatarUrl    string   `json:"avatar_url"`
	Nickname     string   `json:"nickname"`
	IsAdmin      bool     `json:"is_admin"`
	FlagSum      int      `json:"flag_sum"`
	UserIntreTag []string `json:"user_intre_tag"`
	FlaggedToday bool     `json:"flagged_today"`
	UserLevel    int      `json:"user_level"`
}

type FindFlagger struct {
	Fid           int                      `json:"fid"`
	FlaggerTitle  string                   `json:"flagger_title"`
	TagTitle      string                   `json:"tag_title"`
	Announcement  string                   `json:"announcement"`
	IsMember      bool                     `json:"is_member"`
	ShouldFlagSum int                      `json:"should_flag_sum"`
	FlaggerMember []FlaggerGroupMemberInfo `json:"flagger_member"`
}

type FlaggerGroupMemberInfoPlus struct {
	Uid                int      `json:"uid"`
	AvatarUrl          string   `json:"avatar_url"`
	Nickname           string   `json:"nickname"`
	IsAdmin            bool     `json:"is_admin"`
	FlagSum            int      `json:"flag_sum"`
	UserIntreTag       []string `json:"user_intre_tag"`
	SequentialFlagTime int      `json:"sequential_flag_time"`
	FlaggedToday       bool     `json:"flagged_today"`
	UserLevel          int      `json:"user_level"`
}

type FlaggerInfo struct {
	Fid           int                          `json:"fid"`
	FlaggerTitle  string                       `json:"flagger_title"`
	TagTitle      string                       `json:"tag_title"`
	Announcement  string                       `json:"announcement"`
	IsMember      bool                         `json:"is_member"`
	ShouldFlagSum int                          `json:"should_flag_sum"`
	FlaggerMember []FlaggerGroupMemberInfoPlus `json:"flagger_member"`
}

type ReturnTagsInfo struct {
	UserIntreTag []string `json:"user_intre_tag"`
	AllTags      []string `json:"all_tags"`
}

type UserInfo struct {
	Nickname        string   `json:"nickname"`
	AvatarUrl       string   `json:"avatar_url"`
	Grade           int      `json:"grade"`
	Major           int      `json:"major"`
	UserSocialTrend []int    `json:"user_social_trend"`
	Environment     []int    `json:"environment"`
	UserIntreTag    []string `json:"user_intre_tag"`
	HaveFlaged      int      `json:"have_flaged"`
	ShouldFlagSum   int      `json:"should_flag_sum"`
	CredenceValue   int      `json:"credence_value"`
	// UserHistory     []FindFlagger
}
