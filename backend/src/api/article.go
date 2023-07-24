// Package api article 相关接口
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
	"time"
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
	//	Keyword    string `json:"keyword"`
	//	Likes      string `json:"likes"`
	//	State      int    `json:"state"`
	pageNum, err := utils.StrToInt(c.Query("page_num"))
	if err != nil {
		pageNum = PageNum
	}
	pageSize, err := utils.StrToInt(c.Query("page_size"))
	if err != nil {
		pageSize = PageSize
	}
	TagID, err := utils.StrToInt(c.Query("tag_id"))
	if err != nil {
		TagID = 0
	}
	categoryID, err := utils.StrToInt(c.Query("category_id"))
	if err != nil {
		categoryID = 0
	}

	query := conn
	if TagID != 0 {
		query = query.Joins("left join article_tag_rel on articles.id = article_tag_rel.article_id").
			Where("article_tag_rel.tag_id = ?", TagID)
	}
	if categoryID != 0 {
		query = query.Joins("left join article_category_rel on articles.id = article_category_rel.article_id").
			Where("article_category_rel.article_category_id = ?", categoryID)
	}

	res := query.Scopes(models.Paginate(pageNum, pageSize)).Preload("Meta").Find(&articles)

	http.SuccessResponse(c, &map[string]any{
		"list":  articles,
		"count": res.RowsAffected,
	})
}

func GetArchiveList(c *gin.Context) {
	var articles []models.Article
	conn := models.GetDBConn()

	conn.Select("id", "created_at", "title").Order("created_at").Find(&articles)
	articleDic := make(map[int][]models.Article)
	yearCount := 0
	for _, article := range articles {
		tm := time.Unix(0, article.CreatedAt)
		year := tm.Year()
		_, ok := articleDic[year]
		if !ok {
			yearCount++
		}
		articleTmp := make(map[string]any)
		articleTmp["id"] = article.ID
		articleTmp["Title"] = article.Title
		articleTmp["created_at"] = article.CreatedAt
		articleDic[year] = append(articleDic[year], article)
	}

	var articleList []map[string]any
	for k, v := range articleDic {
		tmp := make(map[string]any)
		tmp["year"] = utils.IntToStr(k)
		tmp["list"] = v
		articleList = append(articleList, tmp)
	}

	http.SuccessResponse(c, &map[string]any{
		"list":  articleList,
		"count": yearCount,
	})
}

func GetArticleDetail(c *gin.Context) {
	conn := models.GetDBConn()
	articleType, err := utils.StrToInt(c.Query("type"))
	if err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	var article = models.Article{}
	if articleType == 3 {
		article.ArticleType = 3
		conn.Where(&article).First(&article)

	} else {
		articleID, err := utils.StrToUInt64(c.Query("id"))
		if err != nil {
			http.ErrorResponse(c, err.Error())
			return
		}
		article.ID = articleID
		conn.Preload("Meta").Preload("Tags").First(&article)
		var comments []models.Comment
		conn.Preload("User").Preload("OtherComments").
			Where("article_id = ? and comment_id = ?", article.ID, 0).Find(&comments)
		article.Comment = comments
	}
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
		return
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
	pageNum, err := utils.StrToInt(c.Query("page_num"))
	if err != nil {
		pageNum = PageNum
	}
	pageSize, err := utils.StrToInt(c.Query("page_size"))
	if err != nil {
		pageSize = PageSize
	}
	conn := models.GetDBConn()
	var tags []models.Tag

	res := conn.Model(&models.Tag{}).Scopes(models.Paginate(pageNum, pageSize)).Find(&tags)
	http.SuccessResponse(c, &map[string]any{
		"list":  tags,
		"count": res.RowsAffected,
	})
}
