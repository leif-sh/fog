package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/docs"
	"github.com/leif-sh/fog/src/api"
	"github.com/leif-sh/fog/src/middleware"
	"github.com/leif-sh/fog/src/models"
	"github.com/leif-sh/fog/src/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func setupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.Logger())
	docs.SwaggerInfo.BasePath = "/api"
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/register", api.Register)
		apiGroup.POST("/login", api.Login)
		apiGroup.POST("/logout", api.Logout)

		apiGroup.GET("/getArchiveList", api.GetArchiveList)
		apiGroup.GET("/getArticleList", api.GetArticleList)
		apiGroup.GET("/getArticleDetail", api.GetArticleDetail)
		apiGroup.POST("/likeArticle", api.LikeArticle)

		apiGroup.POST("/addComment", api.AddComment)
		apiGroup.POST("/addThirdComment", api.AddThirdComment)
		apiGroup.GET("/getCommentList", api.GetCommentList)
		apiGroup.GET("/getTagList", api.GetTagList)

		apiGroup.GET("/getProjectList", api.GetProjectList)
		apiGroup.GET("/getProjectDetail", api.GetProjectDetail)

		apiGroup.GET("/getTimeAxisList", api.GetTimeLineList)
		apiGroup.GET("/getTimeAxisDetail", api.GetTimeLineDetail)

	}

	// 注册Swagger api相关路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Authorized group (uses gin.BasicAuth() middleware)
	adminGroup := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	var db = make(map[string]string)
	adminGroup.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return router
}

func main() {
	models.InitDB()
	utils.InitLog()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	gin.SetMode(gin.DebugMode)
	err := r.Run(":8001")
	if err != nil {
		return
	}
}
