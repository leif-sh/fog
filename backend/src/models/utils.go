package models

import (
	"gorm.io/gorm"
)

/*
*
分页封装
*/
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

/*
*
排序判断
*/
func Order(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 判断字符串第一个字符是 + -
		mark := sort[0:1]
		orderStr := sort[1:]
		switch {
		case mark == "-":
			orderStr = orderStr + ` ASC`
		default:
			orderStr = orderStr + ` DESC`
		}
		return db.Order(orderStr)
	}
}

/*
*
where查询
*/
func Where(wheres map[string]map[string]string) func(db *gorm.DB) *gorm.DB {
	var andWhere, orWhere = map[string]interface{}{}, map[string]interface{}{}
	var likeWhere string
	for key, value := range wheres {
		//组装where
		if key == "AND" {
			for key2, value2 := range value {
				andWhere[key2] = value2
			}
		}
		if key == "OR" {
			for key2, value2 := range value {
				orWhere[key2] = value2
			}
		}
		if key == "LIKE" {
			for key2, value2 := range value {
				likeWhere += key2 + ` LIKE "` + value2 + `%" OR `
			}
			// 去除尾部OR
			likeWhere = likeWhere[:len(likeWhere)-3]
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(andWhere).Where(likeWhere).Or(orWhere)
	}
}
