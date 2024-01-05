package router

import (
	"go-back/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(dbConn *gorm.DB) {

	articleHandler := controller.ArticleHandler{
		Db: dbConn,
	}

	r := gin.Default()
	setCors(r)

	// r.LoadHTMLGlob("view/*html")
	r.GET("/articles", articleHandler.ShowAllArticle)
	// r.GET("/show/:id", ShowOneBlog)
	r.Run(":9001")

}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}
