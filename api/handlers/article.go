package handlers

import (
	"fmt"

	"github.com/Dovudxon2004/Api/Article/models"

	"github.com/gin-gonic/gin"
)

// Create Article godoc
// @tags article
//  @ID create-article-handler
// @Summary Create article
// @Description Create article by gicen info and ID
// @Param data body models.ArticleCreateModel true "Article body"
// @Accept json
// @Produce json
// @Succes 200 {object} models.SUccesResponse{data=string}
// @Failure default {objext} models.DefaultError
// @Router /articles [POST]
func (h *Handler) CreateArticle(c *gin.Context) {
	var entity models.ArticleCreateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	fmt.Println(entity)

	err = h.strg.Article().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been created",
		Data:    "ok",
	})
}

func (h *Handler) GetArticleList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})

		limit, err := h.getLimitParam(c)
		if err != nil {
			c.JSON(400, models.DefaultError{
				Message: err.Error(),
			})
		}

		resp, err := h.strg.Article().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})
		if err != nil {
			c.JSON(400, models.DefaultError{
				Message: err.Error(),
			})
		}

		c.JSON(200, resp)
	}

}
