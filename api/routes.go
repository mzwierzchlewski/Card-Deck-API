package api

import (
	"github.com/gin-gonic/gin"

	"card-deck-api/components/deck/routes"
)

// @title Card Deck API
// @version 0.1
// @description This a card deck api server.
// @description Made for Toggl.

// @contact.name Maciej Zwierzchlewski
// @contact.url https:///maciejz.dev
// @contact.email zwierzchlewski.maciej@outlook.com

// @host localhost:13370
// @BasePath /

func RegisterGinRouter() (router *gin.Engine) {
	router = gin.Default()
	deckroutes.RegisterDeckRoutes(router)
	AddSwagger(router, true)

	return
}
