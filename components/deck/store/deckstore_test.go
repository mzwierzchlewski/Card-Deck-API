package deckstore

import (
	"fmt"
	"testing"

	"card-deck-api/models"
)

func TestGet(t *testing.T) {
	table := []struct {
		input          string
		expectedDeckId interface{}
		expectedExists bool
	}{
		{"1", "1", true},
		{"2", nil, false},
	}

	for _, test := range table {
		testName := fmt.Sprintf("Get(%s)", test.input)
		t.Run(testName, func(t *testing.T) {
			// Arrange
			teardown := setup(t)

			// Act
			resultDeck, resultExists := Get(test.input)

			// Assert
			if test.expectedExists {
				expectedDeckId := test.expectedDeckId.(string)
				if !resultExists || resultDeck == nil || resultDeck.ID != expectedDeckId {
					t.Errorf("Got %#v, %t, expected deck with ID %s and exists: %t", resultDeck, resultExists, expectedDeckId, test.expectedExists)
				}
			} else if resultExists || resultDeck != nil {
				t.Errorf("Got %#v, %t, expected no deck and exists: %t", resultDeck, resultExists, test.expectedExists)
			}

			// Clean-up
			teardown(t)
		})
	}
}

func TestSet_WithNewDeck_SetsDeckInCorrectPlace(t *testing.T) {
	// Arrange
	teardown := setup(t)
	id := "12345"
	newDeck := models.Deck{ID: id}

	// Act
	Set(newDeck)

	// Assert
	if decks[id].ID != id {
		t.Errorf("Deck set in wrong place, decks: %v", decks)
	}

	// Clean-up
	teardown(t)
}

func TestSet_WithNewDeck_DeepCopiesTheDeck(t *testing.T) {
	// Arrange
	teardown := setup(t)
	id, code1, code2 := "12345", "AK", "Foo"
	cards := []models.Card{{Code: code1}}
	newDeck := models.Deck{ID: id, Cards: cards}

	// Act
	Set(newDeck)
	newDeck.Cards[0].Code = code2

	// Assert
	if decks[id].Cards[0].Code != code1 {
		t.Errorf("Deck not deep copied before adding to store, expected first card code: %s, got: %s", code1, code2)
	}

	// Clean-up
	teardown(t)
}

func TestSet_WithExistingDeck_ReplacesExistingDeck(t *testing.T) {
	// Arrange
	teardown := setup(t)
	id, code := "1", "AK"
	cards := []models.Card{{Code: code}}
	newDeck := models.Deck{ID: id, Cards: cards}

	// Act
	Set(newDeck)

	// Assert
	if len(decks[id].Cards) != 1 && decks[id].Cards[0].Code != code {
		t.Errorf("Deck not replaced in memory, expected: %#v, got %#v", newDeck, decks[id])
	}

	// Clean-up
	teardown(t)
}

func TestModify_WithNonExistingDeck_ReturnsNils(t *testing.T) {
	// Arrange
	teardown := setup(t)
	action := func(deck *models.Deck) {
		return
	}
	// Act
	oldDeck, newDeck := Modify("2", action)

	// Assert
	if oldDeck != nil || newDeck != nil {
		t.Errorf("Expected to get nils, got %#v and %#v", oldDeck, newDeck)
	}
	// Clean-up
	teardown(t)
}

func TestModify_WithExistingDeck_SetsDeckProperty(t *testing.T) {
	// Arrange
	teardown := setup(t)
	action := func(deck *models.Deck) {
		deck.Shuffled = true
	}

	// Act
	oldDeck, newDeck := Modify("1", action)

	// Assert
	if oldDeck.Shuffled || !newDeck.Shuffled {
		t.Errorf("Expected to get false and true, got %t and %t", oldDeck.Shuffled, newDeck.Shuffled)
	}
	// Clean-up
	teardown(t)
}

func TestModify_WhenChangingDecksCards_SetsRemainingPropertyToCorrectValue(t *testing.T) {
	// Arrange
	teardown := setup(t)
	action := func(deck *models.Deck) {
		deck.Cards = []models.Card{{}}
		deck.Remaining = 12
	}

	// Act
	_, newDeck := Modify("1", action)

	// Assert
	if newDeck.Remaining != 1 {
		t.Errorf("Expected to get deck with Remaining: 1, got %#v", newDeck)
	}
	// Clean-up
	teardown(t)
}

func TestModify_BeforeSavingDeck_DeepCopiesDeck(t *testing.T) {
	// Arrange
	teardown := setup(t)
	code1, code2 := "AK", "Foo"
	cards := []models.Card{{Code: code1}}
	action := func(deck *models.Deck) {
		deck.Cards = cards
	}

	// Act
	Modify("1", action)
	cards[0].Code = code2

	// Assert
	if decks["1"].Cards[0].Code != code1 {
		t.Errorf("Saved deck was modified outside of store")
	}
	// Clean-up
	teardown(t)
}

func TestModify_WhenDeckIdIsChanged_Panics(t *testing.T) {
	// Arrange
	teardown := setup(t)
	action := func(deck *models.Deck) {
		deck.ID = "2"
	}
	defer func() {
		// Assert
		if r := recover(); r == nil {
			t.Errorf("Expected panic, code did not panic")
		}
		// Clean-up
		teardown(t)
	}()

	// Act
	Modify("1", action)
}

func setup(tb testing.TB) func(tb testing.TB) {
	decks["1"] = models.Deck{ID: "1"}

	return func(tb testing.TB) {
		decks = make(map[string]models.Deck)
	}
}
