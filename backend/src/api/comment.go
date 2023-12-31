package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
)

func GetCommentList(c *gin.Context) {
	var comments []models.Comment
	conn := models.GetDBConn()
	pageSize, err := utils.StrToInt(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	pageNum, err := utils.StrToInt(c.DefaultQuery("pageSize", "1"))
	if pageNum < 1 {
		http.ErrorResponse(c, "page num begin from 1")
	}
	if err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	res := conn.Model(models.Comment{}).Preload("User").Preload("OtherComments").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	if res.Error != nil {
		http.ErrorResponse(c, res.Error.Error())
		return
	}
	http.SuccessResponse(c, &map[string]any{
		"list":  comments,
		"count": res.RowsAffected,
	})
}

func AddComment(c *gin.Context) {
	conn := models.GetDBConn()
	newComment := models.Comment{}
	if err := c.ShouldBind(&newComment); err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	conn.Create(&newComment)
	var meta models.Meta
	conn.Where("article_id = ?", newComment.ArticleID).First(&meta)
	meta.Comments++
	conn.Save(&meta)
	http.SuccessResponse(c, "success")
}

func AddThirdComment(c *gin.Context) {
	conn := models.GetDBConn()
	newComment := models.Comment{}

	if err := c.ShouldBind(&newComment); err != nil {
		utils.SugarLogger.Error(err.Error())
		http.ErrorResponse(c, err.Error())
		return
	}
	conn.Create(&newComment)
	http.SuccessResponse(c, "success")
}
