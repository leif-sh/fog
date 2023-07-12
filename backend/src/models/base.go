package models

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type BaseModel struct {
	ID        uint64 `gorm:"primaryKey;column:id" json:"id"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"` // 使用时间戳毫秒数填充更新时间
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"` // 使用时间戳秒数填充创建时间
}

// Scan 解码json字符串
type JsonUtils string

func (card *JsonUtils) Scan(val interface{}) error {
	b, _ := val.([]byte)
	return json.Unmarshal(b, card)
}

// Value 编码json
func (card JsonUtils) Value() (value driver.Value, err error) {
	return json.Marshal(card)
}

type StrList []string

func (m *StrList) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), ",")
	*m = ss
	return nil
}
func (m StrList) Value() (driver.Value, error) {
	str := strings.Join(m, ",")
	return str, nil
}
