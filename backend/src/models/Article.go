package models

type Article struct {
	BaseModel
	CreateTime  int                `json:"create_time"`
	Title       string             `gorm:"size:256;" json:"title"`
	Author      string             `gorm:"size:256;" json:"author"`
	Desc        string             `gorm:"size:256" json:"desc"`
	ImgUrl      string             `gorm:"size:256" json:"img_url"`
	Tags        []*Tag             `gorm:"many2many:article_tag_rel;" json:"tags"`
	Categories  []*ArticleCategory `gorm:"many2many:article_category_rel;" json:"category"`
	Meta        Meta               `gorm:"foreignKey:ArticleID" json:"meta"`
	Comment     []Comment          `gorm:"foreignKey:ArticleID" json:"comments"`
	Keyword     StrList            `gorm:"size:256" json:"keyword"`
	ArticleType int                `json:"type"`    //1: 普通文章，2: 简历，3: 管理员介绍
	State       int                `json:"state"`   // 文章发布状态 => 0 草稿，1 已发布
	Origin      int                `json:"origin"`  // 文章转载状态 => 0 原创，1 转载，2 混合
	Numbers     int                `json:"numbers"` //字数
	Content     string             `gorm:"type:text;" json:"content"`
	LikeUsers   []*User            `gorm:"many2many:article_like_user_rel;" json:"like_users"`
}

type ArticleCategory struct {
	BaseModel
	Name     string     `gorm:"size:256"`
	Desc     string     `gorm:"size:256" json:"desc"`
	Articles []*Article `gorm:"many2many:article_category_rel;"`
}

type Meta struct {
	BaseModel
	ArticleID uint64 `json:"-"`
	Views     int    `json:"views"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
}
