package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
)

func GetProjectList(c *gin.Context) {
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
