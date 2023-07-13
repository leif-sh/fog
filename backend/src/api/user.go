package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/http"
	"github.com/leif-sh/fog/src/models"
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
	var newUser models.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		return
	}
	conn.Create(&newUser)
	http.SuccessResponse(c, "success")
}

func Logout(c *gin.Context) {
	conn := models.GetDBConn()
	var newUser models.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		return
	}
	conn.Create(&newUser)
	http.SuccessResponse(c, "success")
}
