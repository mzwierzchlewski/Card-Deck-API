package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "card-deck-api/docs"
)

func AddSwagger(router *gin.Engine, setSwaggerAsDefault bool) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if setSwaggerAsDefault {
		router.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
	}
}
