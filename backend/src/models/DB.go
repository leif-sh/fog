package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
	"time"
)

var db *gorm.DB

func InitDB() {
	fmt.Println("初始化数据库链接")
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:fog@blog@tcp(localhost:3306)/fog?charset=utf8&loc=Local&parseTime=true",
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("can't connect to mysql")
		os.Exit(-1)
	}
	// 自动更新表结构
	err = db.AutoMigrate(&Article{}, &Meta{}, &Tag{}, &ArticleCategory{}, &Comment{}, &User{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	if err != nil {
		panic(err.Error())
	}
}
func GetDBConn() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
