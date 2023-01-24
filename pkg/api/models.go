package api

import "time"

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password,omitempty"`
}

type ChatMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Time     time.Time
}
