package deckstore

import (
	"sync"

	. "card-deck-api/models"
)

var lock sync.RWMutex
var decks map[string]Deck

func Get(id string) (deck *Deck, exists bool) {
	lock.RLock()
	defer lock.RUnlock()
	sourceDeck, exists := decks[id]
	if !exists {
		return nil, false
	}
	return deepCopyDeck(sourceDeck), true
}

func Modify(id string, action func(*Deck)) (old *Deck, new *Deck) {
	lock.Lock()
	defer lock.Unlock()
	deck, exists := decks[id]
	if !exists {
		return nil, nil
	}
	oldDeck, newDeck := deepCopyDeck(deck), deepCopyDeck(deck)
	action(newDeck)
	if newDeck.ID != id {
		panic("Deck ID illegally changed")
	}
	newDeck.Remaining = len(newDeck.Cards)
	decks[id] = *deepCopyDeck(*newDeck)
	return oldDeck, newDeck
}

func Set(deck Deck) {
	lock.Lock()
	defer lock.Unlock()
	copiedDeck := deepCopyDeck(deck)
	decks[deck.ID] = *copiedDeck
}

func deepCopyDeck(deck Deck) *Deck {
	newDeck := deck
	newDeck.Cards = make([]Card, len(deck.Cards))
	copy(newDeck.Cards, deck.Cards)
	return &newDeck
}

func init() {
	lock = sync.RWMutex{}
	decks = make(map[string]Deck)
}
