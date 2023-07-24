package models

type Message struct {
	BaseModel
	UserID  uint64 `json:"user_id"`
	User    User   `json:"user"`
	Content string `gorm:"type:text" json:"content"`

	State int `json:"state"`
}
