package models

type UserDoingFlagger struct {
	FlaggerTitle      string   `json:"flagger_title"`
	FinishedAvatarUrl []string `json:"finished_avatar_url"`
	FinishedNum       int      `json:"finished_num"`
	FlaggerProgress   string   `json:"flagger_progress"`
}

