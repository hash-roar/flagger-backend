package models

type UserDoingFlagger struct {
	FlaggerTitle      string   `json:"flagger_title,omitempty"`
	FinishedAvatarUrl []string `json:"finished_avatar_url,omitempty"`
	FinishedNum       int      `json:"finished_num,omitempty"`
	FlaggerProgress   string   `json:"flagger_progress,omitempty"`
}
