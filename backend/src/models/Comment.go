package models

type Comment struct {
	BaseModel
	ArticleID     uint64    `json:"article_id"`
	UserID        uint64    `json:"user_id"`
	User          User      `json:"user"`
	ToUserID      uint64    `json:"to_user_id"`
	ToUser        User      `json:"to_user"`
	Content       string    `gorm:"type:text;" json:"content"`
	IsTop         bool      `gorm:"type:tinyint" json:"is_top"`
	IsHandle      bool      `gorm:"type:tinyint" json:"is_handle"`
	Likes         int       `json:"likes"`
	CommentID     uint64    `json:"comment_id"`
	OtherComments []Comment `gorm:"foreignKey:CommentID" json:"other_comments"`
}
