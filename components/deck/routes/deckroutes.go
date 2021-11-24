package deckroutes

import (
	"github.com/gin-gonic/gin"

	"card-deck-api/components/deck/controllers"
)

func RegisterDeckRoutes(router *gin.Engine) {
	router.GET("/decks/:id", deckcontroller.GetDeck)
	router.POST("/decks", deckcontroller.PostDeck)
	router.PATCH("/decks/:id/:numberOfCards", deckcontroller.DrawCards)
}
