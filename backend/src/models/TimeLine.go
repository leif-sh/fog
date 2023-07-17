package models

type TimeLine struct {
	BaseModel
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Content   string `gorm:"type:text" json:"content"`
	Title     string `gorm:"size:256" json:"title"`
	State     int    `json:"state"`
}
