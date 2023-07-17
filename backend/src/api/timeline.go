package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
)

func GetTimeLineList(c *gin.Context) {
	conn := models.GetDBConn()
	pageSize, err := utils.StrToInt(c.Query("page_size"))
	if err != nil {
		pageSize = PageSize
	}
	pageNum, err := utils.StrToInt(c.Query("page_num"))
	if err != nil {
		pageNum = PageNum
	}
	var timelines []models.TimeLine
	res := conn.Scopes(models.Paginate(pageNum, pageSize)).Find(&timelines)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	http.SuccessResponse(c, &map[string]any{
		"list":  timelines,
		"count": res.RowsAffected,
	})
}

func GetTimeLineDetail(c *gin.Context) {
	conn := models.GetDBConn()
	timelineID, err := utils.StrToUInt64(c.Query("id"))
	if err != nil {
		http.ErrorResponse(c, err.Error())
		return
	}
	timeline := models.TimeLine{
		BaseModel: models.BaseModel{
			ID: timelineID,
		},
	}
	conn.First(&timeline)
	http.SuccessResponse(c, timeline)

}
