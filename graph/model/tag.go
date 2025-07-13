package model

type Tag struct {
	ID      string `json:"id"`
	TagName string `json:"tagName"`
	UserID  string `json:"userId"`
	User    *User  `json:"user"`
}
