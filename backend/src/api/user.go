package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
)

func Register(c *gin.Context) {
	conn := models.GetDBConn()
	var newUser models.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		return
	}
	conn.Create(&newUser)
	http.SuccessResponse(c, "success")
}

func Login(c *gin.Context) {
	conn := models.GetDBConn()
	var User models.User
	var rBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&rBody); err != nil {
		utils.SugarLogger.Error("error param")
		http.ErrorResponse(c, err.Error())
		return
	}
	password := utils.MD5(rBody.Password)
	res := conn.Where("email = ? and password = ?", rBody.Email, password).First(&User)
	if res.RowsAffected == 0 {
		http.ErrorResponse(c, "user not exist")
		return
	}
	http.SuccessResponse(c, models.User{
		BaseModel: models.BaseModel{
			ID: User.ID,
		},
		Name:   User.Name,
		Avatar: User.Avatar,
	})
}

func Logout(c *gin.Context) {
	http.SuccessResponse(c, "success")
}
