package models

type Post struct {
	Id     uint   `json:"id,omitempty"`
	UserId uint   `json:"user_id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	User   User   `gorm:"foreignkey:UserId" json:"user"`
}
