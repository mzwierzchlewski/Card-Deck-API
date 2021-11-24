package deckservice

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"card-deck-api/components/cardshuffle"
	"card-deck-api/components/deck/store"
	. "card-deck-api/models"
)

func Create(shuffled bool, cardList []string) (deck Deck, invalidCodes []string) {
	shouldSetOrder := len(cardList) == 0
	if cardList == nil || len(cardList) == 0 {
		cardList = GetPossibleCardCodes()
	}

	cards, invalidCodes := createCards(cardList)
	if shuffled {
		cards = cardshuffleservice.ShuffleCards(cards)
	} else if shouldSetOrder {
		cards = cardshuffleservice.OrderCards(cards)
	}

	deck = Deck{
		ID:        uuid.NewString(),
		Shuffled:  shuffled,
		Remaining: len(cards),
		Cards:     cards,
	}
	deckstore.Set(deck)
	return
}

func Get(id string) (*Deck, error) {
	deck, exists := deckstore.Get(id)
	if !exists {
		errorMsg := fmt.Sprintf("Deck with id %s does not exist.", id)
		return nil, errors.New(errorMsg)
	}
	return deck, nil
}

func Draw(id string, numberOfCards int) ([]Card, error) {
	var drawnCards = make([]Card, 0, numberOfCards)
	oldDeck, _ := deckstore.Modify(id, func(deck *Deck) {
		if numberOfCards > deck.Remaining {
			numberOfCards = deck.Remaining
		}
		drawnCards = deck.Cards[:numberOfCards]
		deck.Cards = deck.Cards[numberOfCards:]
	})
	if oldDeck == nil {
		errorMsg := fmt.Sprintf("Deck with id %s does not exist.", id)
		return nil, errors.New(errorMsg)
	}

	return drawnCards, nil
}

func createCards(cardList []string) (cards []Card, invalidCodes []string) {
	for _, cardCode := range cardList {
		newCard, err := NewCard(cardCode)
		if err != nil {
			invalidCodes = append(invalidCodes, cardCode)
			continue
		}
		cards = append(cards, *newCard)
	}
	return
}
