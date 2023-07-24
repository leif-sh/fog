package models

type Tag struct {
	BaseModel
	Articles []*Article `gorm:"many2many:article_tag_rel;"`
	Name     string     `json:"name"`
	Desc     string     `gorm:"size:256"  json:"desc"`
	Icon     string     `gorm:"size:256" json:"icon"`
}
