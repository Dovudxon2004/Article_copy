package api

import (
	// "github.com/Dovudxon2004/Api/Article/docs"
	// _ "bootcamp/article/api/docs"
	"github.com/Dovudxon2004/Api/Article/api/handlers"
	"github.com/Dovudxon2004/Api/Article/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUpApi(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	r.POST("/articles", h.CreateArticle)
	r.GET("/articles", h.GetArticleList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
