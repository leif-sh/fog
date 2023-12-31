package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
)

func GetProjectList(c *gin.Context) {
	conn := models.GetDBConn()
	pageSize, err := utils.StrToInt(c.Query("page_size"))
	if err != nil {
		pageSize = PageSize
	}
	pageNum, err := utils.StrToInt(c.Query("page_num"))
	if err != nil {
		pageNum = PageNum
	}
	var projects []models.Project
	res := conn.Scopes(models.Paginate(pageNum, pageSize)).Find(&projects)
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	http.SuccessResponse(c, &map[string]any{
		"list":  projects,
		"count": res.RowsAffected,
	})
}

func GetProjectDetail(c *gin.Context) {
	conn := models.GetDBConn()
	projectID, err := utils.StrToUInt64(c.Query("id"))
	if err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	project := models.Project{
		BaseModel: models.BaseModel{
			ID: projectID,
		},
	}
	conn.First(&project)
	http.SuccessResponse(c, project)
}
