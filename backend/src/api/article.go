package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// GetArticleList
// @Summary 摘要
// @Schemes
// @Description 描述
// @Tags 标签
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any 返回结果 200 类型（object就是结构体） 类型 注释
// @Router /api/getArticleList [get]
func GetArticleList(c *gin.Context) {
	var articles []models.Article

	res := conn.Model(&models.Article{}).Limit(20).Find(&articles)
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	utils.SuccessResponse(c, &map[string]any{
		"list":  articles,
		"count": res.RowsAffected,
	})
}

// GetTagList
// @Summary 摘要
// @Schemes
// @Description 描述
// @Tags 标签
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any 返回结果 200 类型（object就是结构体） 类型 注释
// @Router /api/getTagList [get]
func GetTagList(c *gin.Context) {
	var tags []models.Tag

	res := conn.Model(&models.Tag{}).Limit(20).Find(&tags)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	utils.SuccessResponse(c, &map[string]any{
		"list":  tags,
		"count": res.RowsAffected,
	})
}

var conn *gorm.DB

func init() {
	fmt.Println("初始化数据库链接")
	var err error
	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:fog@blog@tcp(localhost:3306)/fog?charset=utf8&loc=Local&parseTime=true",
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("can't connect to mysql")
		os.Exit(-1)
	}
	// 自动更新表结构
	err = conn.AutoMigrate(&models.Article{}, &models.Meta{}, &models.Tag{}, &models.ArticleCategory{})
	if err != nil {
		fmt.Println(err)
	}
}
