package models

type User struct {
	BaseModel
	Email     string `gorm:"size:256;unique:true;" json:"email"`
	Name      string `gorm:"size:256;default:" json:"name"`
	Password  string `gorm:"size:256;" json:"password"`
	Phone     string `gorm:"size:20;" json:"phone"`
	Desc      string `gorm:"size:256;" json:"desc"`
	Avatar    string `gorm:"size:256" json:"avatar"`
	ImgUrl    string `gorm:"size:256" json:"img_url"`
	Introduce string `gorm:"size:256" json:"introduce"`
	Location  string `gorm:"size:256" json:"location"`
	UserType  int    `json:"type"` //用户类型 0：博主，1：其他用户 ，2：github， 3：weixin， 4：qq
}

type APIUser struct {
	Email  string `gorm:"size:256;" json:"email"`
	Name   string `gorm:"size:256;default:" json:"name"`
	Phone  string `gorm:"size:20;" json:"phone"`
	Desc   string `gorm:"size:256;" json:"desc"`
	Avatar string `gorm:"size:256" json:"avatar"`
}
