package models

type Comment struct {
	BaseModel
	ArticleID     uint64    `json:"article_id"`
	UserID        uint64    `json:"user_id"`
	User          User      `json:"user"`
	ToUserID      uint64    `json:"to_user_id"`
	Content       string    `gorm:"type:text;" json:"content"`
	IsTop         bool      `gorm:"type:tinyint" json:"is_top"`
	IsHandle      bool      `gorm:"type:tinyint" json:"is_handle"`
	Likes         int       `json:"likes"`
	OtherComments []Comment `json:"other_comments"`
}
