package models

type User struct {
	BaseModel
	Email    string `gorm:"size:256;" json:"email"`
	Name     string `gorm:"size:256;default:" json:"name"`
	Password string `gorm:"size:256;" json:"password"`
	Phone    string `gorm:"size:20;" json:"phone"`
	Desc     string `gorm:"size:256;" json:"desc"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	UserType int    `json:"type"`
}
