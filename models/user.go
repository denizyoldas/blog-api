package models

type User struct {
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `gorm:"unique" json:"email"`
	Password    []byte `gorm:"->;<-;" json:"password"`
	Posts       []Post `json:"posts,omitempty"`
}
