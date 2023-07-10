package models

type BaseModel struct {
	ID        uint  `gorm:"primaryKey;column:id" json:"id"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"` // 使用时间戳毫秒数填充更新时间
	CreatedAt int64 `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"` // 使用时间戳秒数填充创建时间
}
