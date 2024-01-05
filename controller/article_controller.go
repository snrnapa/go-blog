package controller

import (
	"go-back/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleHandler struct {
	Db *gorm.DB
}

func (a *ArticleHandler) ShowAllArticle(c *gin.Context) {
	var articles []entity.Article
	a.Db.Find(&articles)
	c.JSON(http.StatusOK, articles)

}
