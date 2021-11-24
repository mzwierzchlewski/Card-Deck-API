package deckcontroller

import . "card-deck-api/models"

type apiError struct {
	Error string `json:"error" example:"Deck with id e89efc4d-9294-443a-8a54-d78cd0e8a0c9 does not exist."`
}

type postDeckRequest struct {
	Shuffled bool     `json:"shuffled" example:"false"`
	Cards    []string `json:"cards" example:"AS,KD,AC,2C,KH"`
}

type postDeckResponse struct {
	ID           string   `json:"deck_id" example:"e89efc4d-9294-443a-8a54-d78cd0e8a0c9"`
	Shuffled     bool     `json:"shuffled" example:"false"`
	Remaining    int      `json:"remaining" example:"52"`
	InvalidCards []string `json:"invalid_cards" example:"0H,2P"`
}

type drawCardsResponse struct {
	Cards []Card `json:"cards"`
}
