// Package api article 相关接口
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"strconv"
)

// GetArticleList
// @Summary 获取文章列表
// @Schemes
// @Description 获取文章列表
// @Tags article
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any 返回结果 200 类型（object就是结构体） 类型 注释
// @Router /api/getArticleList [get]
func GetArticleList(c *gin.Context) {
	var articles []models.Article
	conn := models.GetDBConn()
	res := conn.Preload("Meta").Model(&models.Article{}).Limit(20).Find(&articles)
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	http.SuccessResponse(c, &map[string]any{
		"list":  articles,
		"count": res.RowsAffected,
	})
}

func GetArticleDetail(c *gin.Context) {
	conn := models.GetDBConn()
	articleID := c.Query("id")
	_id, err := strconv.Atoi(articleID)
	if err != nil {
		http.ErrorResponse(c, err.Error())
	}
	var article = models.Article{
		BaseModel: models.BaseModel{
			ID: uint64(_id),
		},
	}
	res := conn.Preload("Meta").Preload("Comment").Preload("Tags").First(&article)
	if res.Error != nil {
		fmt.Println(res.Error)
		http.ErrorResponse(c, res.Error.Error())
	}
	http.SuccessResponse(c, article)
}

func AddComment(c *gin.Context) {
	conn := models.GetDBConn()
	newComment := models.Comment{}
	if err := c.ShouldBind(&newComment); err != nil {
		http.ErrorResponse(c, err.Error())
	}
	conn.Create(&newComment)
	http.SuccessResponse(c, "success")
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
	conn := models.GetDBConn()
	var tags []models.Tag

	res := conn.Model(&models.Tag{}).Limit(20).Find(&tags)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	http.SuccessResponse(c, &map[string]any{
		"list":  tags,
		"count": res.RowsAffected,
	})
}
