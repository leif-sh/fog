package models

type Comment struct {
	BaseModel
	ArticleID uint64 `json:"article_id"`
	UserID    uint64 `json:"user_id"`
	Content   string `gorm:"size:512;" json:"content"`
}
