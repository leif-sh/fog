package models

type TimeLine struct {
	BaseModel
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Content   string `gorm:"type:text" json:"content"`
	Title     string `gorm:"size:256" json:"title"`
	State     int    `json:"state"` // 状态 1 是已经完成 ，2 是正在进行，3 是没完成
}
