package models

type User struct {
	BaseModel
	Email string `gorm:"size:256;" json:"email"`
}
