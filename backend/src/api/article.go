// Package api article 相关接口
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
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
	articleID, err := utils.StrToInt(c.Query("id"))
	if err != nil {
		http.ErrorResponse(c, err.Error())
	}
	var article = models.Article{
		BaseModel: models.BaseModel{
			ID: uint64(articleID),
		},
	}
	conn.Preload("Meta").Preload("Comment.OtherComments").Preload("Tags").First(&article)
	var comments []models.Comment
	conn.Preload("User").Where("article_id = ?", article.ID).Find(&comments)
	article.Comment = comments
	article.Meta.Views++
	conn.Save(&article.Meta)
	http.SuccessResponse(c, article)
}

func LikeArticle(c *gin.Context) {
	conn := models.GetDBConn()
	var requestBody struct {
		ID     uint64 `json:"id"`
		UserID uint64 `json:"user_id"`
	}
	if err := c.ShouldBind(&requestBody); err != nil {
		utils.SugarLogger.Error("error param", "article_id", requestBody.ID)
		http.ErrorResponse(c, err.Error())
	}
	meta := models.Meta{
		ArticleID: requestBody.ID,
	}
	conn.First(&meta)
	meta.Likes++
	conn.Save(&meta)
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
