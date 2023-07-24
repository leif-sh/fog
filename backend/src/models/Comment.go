package models

type Comment struct {
	BaseModel
	ArticleID     uint64    `json:"article_id"`
	UserID        uint64    `json:"user_id"` //评论人
	User          User      `json:"user"`
	ToUserID      uint64    `json:"to_user_id"` //回复评论
	ToUser        User      `json:"to_user"`
	Content       string    `gorm:"type:text;" json:"content"`
	IsTop         bool      `gorm:"type:tinyint" json:"is_top"`
	IsHandle      bool      `gorm:"type:tinyint" json:"is_handle"` // 是否已经处理过 => 1 是 / 2 否 ；新加的评论需要审核，防止用户添加 垃圾评论
	Likes         int       `json:"likes"`
	CommentID     uint64    `json:"comment_id"` // 回复评论id
	OtherComments []Comment `gorm:"foreignKey:CommentID" json:"other_comments"`
	State         int       `json:"state"` //状态 => 0 待审核 / 1 通过正常 / -1 已删除 / -2 垃圾评论
}
