package models

type Article struct {
	BaseModel
	CreateTime  int                `json:"create_time"`
	Title       string             `gorm:"size:256;" json:"title"`
	Desc        string             `gorm:"size:256" json:"desc"`
	ImgUrl      string             `gorm:"size:256" json:"img_url"`
	Tags        []*Tag             `gorm:"many2many:article_tag_rel;" json:"tags"`
	Categories  []*ArticleCategory `gorm:"many2many:article_category_rel;" json:"category"`
	Meta        Meta               `gorm:"foreignKey:ArticleID" json:"meta"`
	Comment     []Comment          `gorm:"foreignKey:ArticleID" json:"comment"`
	Keyword     StrList            `gorm:"size:256" json:"keyword"`
	ArticleType int                `json:"type"`
	Content     string             `gorm:"type:text;" json:"content"`
}

type ArticleCategory struct {
	BaseModel
	Name     string     `gorm:"size:256"`
	Articles []*Article `gorm:"many2many:article_category_rel;"`
}

type Meta struct {
	BaseModel
	ArticleID uint64 `json:"-"`
	Views     int    `json:"views"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
}

type Tag struct {
	BaseModel
	Articles []*Article `gorm:"many2many:article_tag_rel;"`
	Name     string
}
