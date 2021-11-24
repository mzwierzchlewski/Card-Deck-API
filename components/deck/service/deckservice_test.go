package deckservice

import (
	"testing"

	"card-deck-api/components/deck/store"
	"card-deck-api/models"
)

func TestGet_CorrectDeckId_ReturnsExistingDeck(t *testing.T) {
	// Arrange
	deck := models.Deck{ID: "TestGet_CorrectDeckId_ReturnsExistingDeck", Shuffled: true}
	deckstore.Set(deck)

	// Act
	resultDeck, err := Get("TestGet_CorrectDeckId_ReturnsExistingDeck")

	// Assert
	if resultDeck.ID != "TestGet_CorrectDeckId_ReturnsExistingDeck" || resultDeck.Shuffled != true || err != nil {
		t.Errorf("Got wrong response, expected deck with id = TestGet_CorrectDeckId_ReturnsExistingDeck, shuffled = true and nil error, got: %#v and %v", resultDeck, err)
	}
}

func TestGet_IncorrectDeckId_ReturnsError(t *testing.T) {
	// Arrange

	// Act
	resultDeck, err := Get("TestGet_IncorrectDeckId_ReturnsError")

	// Assert
	if resultDeck != nil || err == nil {
		t.Errorf("Expected nil deck and error, got: %#v and %v", resultDeck, err)
	}
}

func TestDraw_CalledForExistingDeckAndEnoughCards_ReturnsTopCards(t *testing.T) {
	// Arrange
	cards := []models.Card{
		{"King", "Hearts", "KH"},
		{"Ace", "Hearts", "AH"},
		{"2", "Clubs", "2C"},
		{"3", "Spades", "3S"},
		{"3", "Clubs", "3C"},
		{"Ace", "Diamonds", "AD"},
	}
	deck := models.Deck{ID: "TestDraw_CalledForExistingDeckAndEnoughCards_ReturnsTopCards", Shuffled: true, Cards: cards, Remaining: len(cards)}
	deckstore.Set(deck)

	// Act
	drawnCards, err := Draw("TestDraw_CalledForExistingDeckAndEnoughCards_ReturnsTopCards", 3)

	//
	if len(drawnCards) != 3 || err != nil || drawnCards[0].Code != "KH" || drawnCards[1].Code != "AH" || drawnCards[2].Code != "2C" {
		t.Errorf("Expected 3 first cards from %v, got %v, %v", cards, drawnCards, err)
	}
}

func TestDraw_CalledForExistingDeckAndNotEnoughCards_ReturnsAllCards(t *testing.T) {
	// Arrange
	cards := []models.Card{
		{"King", "Hearts", "KH"},
		{"Ace", "Hearts", "AH"},
	}
	deck := models.Deck{ID: "TestDraw_CalledForExistingDeckAndNotEnoughCards_ReturnsAllCards", Shuffled: true, Cards: cards, Remaining: len(cards)}
	deckstore.Set(deck)

	// Act
	drawnCards, err := Draw("TestDraw_CalledForExistingDeckAndNotEnoughCards_ReturnsAllCards", 3)

	//
	if len(drawnCards) != 2 || err != nil || drawnCards[0].Code != "KH" || drawnCards[1].Code != "AH" {
		t.Errorf("Expected 2 first cards from %v, got %v, %v", cards, drawnCards, err)
	}
}

func TestDraw_CalledForNonExistingDeck_ReturnsError(t *testing.T) {
	// Arrange

	// Act
	drawnCards, err := Draw("TestDraw_CalledForNonExistingDeck_ReturnsError", 3)

	// Assert
	if drawnCards != nil || err == nil {
		t.Errorf("Expected nil cards and error got: %v, %v", drawnCards, err)
	}
}

func TestCreate_WithCardsListNotShuffled_ReturnsDeckWithCardsInSameOrder(t *testing.T) {
	// Arrange
	cardsList := []string{"AH", "2C"}

	// Act
	deck, _ := Create(false, cardsList)

	// Assert
	if deck.Shuffled != false || deck.Remaining != 2 || len(deck.Cards) != 2 || deck.Cards[0].Code != "AH" || deck.Cards[1].Code != "2C" {
		t.Errorf("Wrong deck created, expected deck with 2 cards, got: %#v", deck)
	}
}

func TestCreate_WithCardsListWithInvalidCodes_ReturnsDeckWithCardsAndInvalidCodes(t *testing.T) {
	// Arrange
	cardsList := []string{"AH", "TEST"}

	// Act
	deck, invalidCodes := Create(false, cardsList)

	// Assert
	if deck.Shuffled != false || deck.Remaining != 1 || len(deck.Cards) != 1 || deck.Cards[0].Code != "AH" || len(invalidCodes) != 1 || invalidCodes[0] != "TEST" {
		t.Errorf("Wrong deck created, expected deck with 1 card and 1 invalid code, got: %#v", deck)
	}
}

func TestCreate_WithoutCardsList_CreatesWholeDeck(t *testing.T) {
	// Arrange

	// Act
	deck, _ := Create(false, []string{})

	// Assert
	if deck.Shuffled != false || deck.Remaining != 52 || len(deck.Cards) != 52 {
		t.Errorf("Wrong deck created, expected full deck with 52 cards, got: %#v", deck)
	}
}

func TestCreate_WithoutCardsListNotShuffled_ReturnsWholeDeckInOrder(t *testing.T) {
	// Arrange
	correctOrder := []string{
		"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
		"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
		"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
		"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH"}
	// Act
	deck, _ := Create(false, []string{})

	// Assert
	allGood := true
	for i, card := range deck.Cards {
		if card.Code != correctOrder[i] {
			allGood = false
			break
		}
	}

	if !allGood {
		t.Errorf("Cards returned in wrong order, expected: %v, got %#v", correctOrder, deck.Cards)
	}
}
