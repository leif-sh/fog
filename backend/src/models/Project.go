package models

type Project struct {
	BaseModel
	Content   string `gorm:"type:text" json:"content"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Img       string `gorm:"size:256" json:"img"`
	Title     string `gorm:"size:256" json:"title"`
	Url       string `gorm:"size:256" json:"url"`
	State     int    `json:"state"` // 状态 1 是已经完成 ，2 是正在进行，3 是没完成
}
