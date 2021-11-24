package cardshuffleservice

import (
	"testing"

	"card-deck-api/models"
)

func TestOrderCards(t *testing.T) {
	// Arrange
	cards := []models.Card{
		{"King", "Hearts", "KH"},
		{"Ace", "Hearts", "AH"},
		{"2", "Clubs", "2C"},
		{"3", "Spades", "3S"},
		{"3", "Clubs", "3C"},
		{"Ace", "Diamonds", "AD"},
	}
	correctOrder := []string{"2C", "3C", "AD", "AH", "KH", "3S"}

	// Act
	orderedCards := OrderCards(cards)

	// Assert
	for i, card := range orderedCards {
		if card.Code != correctOrder[i] {
			t.Errorf("Wrong order, expected %s at position %d, got %s", correctOrder[i], i, card.Code)
		}
	}
}

func TestShuffleCards(t *testing.T) {
	// Arrange
	cards := []models.Card{
		{"2", "Clubs", "2C"},
		{"3", "Clubs", "3C"},
		{"Ace", "Diamonds", "AD"},
		{"Ace", "Hearts", "AH"},
		{"King", "Hearts", "KH"},
		{"3", "Spades", "3S"},
	}

	// Act
	orderedCards := ShuffleCards(cards)

	// Assert
	for i, card := range orderedCards {
		if card.Code != cards[i].Code {
			t.Errorf("Wrong order, not expected %s at position %d", card.Code, i)
		}
	}
}
