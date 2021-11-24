package deckcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"card-deck-api/components/deck/service"
)

// GetDeck godoc
// @Summary Opens a Deck
// @Description Displays deck and its cards by ID.
// @Tags deck
// @ID get-deck-by-id
// @Produce  json
// @Param id path string true "Deck ID"
// @Success 200 {object} models.Deck
// @Failure 400 {object} apiError
// @Router /decks/{id} [get]
func GetDeck(c *gin.Context) {
	deckId := c.Param("id")
	if len(deckId) == 0 {
		c.JSON(400, apiError{Error: "Invalid/missing deck ID"})
		return
	}

	deck, err := deckservice.Get(deckId)
	if err != nil {
		c.JSON(400, apiError{Error: err.Error()})
		return
	}

	c.JSON(200, deck)
}

// PostDeck godoc
// @Summary Creates a new Deck
// @Description Creates a new deck.
// @Description If no cards are specified, a full deck is created.
// @Description If cards are to be in random order set the shuffled param.
// @Tags deck
// @ID create-deck
// @Accept json
// @Produce  json
// @Param deckOptions body postDeckRequest false "New deck options"
// @Success 201 {object} postDeckResponse
// @Failure 400 {object} apiError
// @Router /decks [post]
func PostDeck(c *gin.Context) {
	request := postDeckRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, apiError{Error: err.Error()})
		return
	}

	deck, invalidCards := deckservice.Create(request.Shuffled, request.Cards)

	response := postDeckResponse{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
	if len(invalidCards) > 0 {
		response.InvalidCards = invalidCards
	}

	c.JSON(201, response)
}

// DrawCards godoc
// @Summary Draws cards from a deck
// @Description Draws cards from the top of the deck.
// @Tags deck
// @ID draw-cards
// @Produce  json
// @Param id path string true "Deck ID"
// @Param numberOfCards path int true "Number of cards to draw"
// @Success 200 {object} drawCardsResponse
// @Failure 400 {object} apiError
// @Router /decks/{id}/{numberOfCards} [patch]
func DrawCards(c *gin.Context) {
	deckId := c.Param("id")
	if len(deckId) == 0 {
		c.JSON(400, apiError{Error: "Invalid deck ID"})
		return
	}
	numberOfCards, err := strconv.Atoi(c.Param("numberOfCards"))
	if numberOfCards == 0 || err != nil {
		c.JSON(400, apiError{Error: "Invalid number of cards"})
		return
	}

	cards, err := deckservice.Draw(deckId, numberOfCards)
	if err != nil {
		c.JSON(400, apiError{Error: err.Error()})
		return
	}

	c.JSON(200, drawCardsResponse{cards})
}
