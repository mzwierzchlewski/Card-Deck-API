package cardshuffleservice

import (
	"math/rand"
	"sort"

	. "card-deck-api/models"
)

func OrderCards(cards []Card) []Card {
	sort.Sort(cardCollection(cards))
	return cards
}

func ShuffleCards(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

type cardCollection []Card

func (c cardCollection) Len() int {
	return len(c)
}

func (c cardCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c cardCollection) Less(i, j int) bool {
	cardISuit, cardIRank := suits[c[i].Suit], ranks[c[i].Value]
	cardJSuit, cardJRank := suits[c[j].Suit], ranks[c[j].Value]

	return cardISuit*100+cardIRank < cardJSuit*100+cardJRank
}

var ranks = map[string]int{
	"2":     1,
	"3":     2,
	"4":     3,
	"5":     4,
	"6":     5,
	"7":     6,
	"8":     7,
	"9":     8,
	"10":    9,
	"Jack":  10,
	"Queen": 11,
	"King":  12,
	"Ace":   0,
}

var suits = map[string]int{
	"Clubs":    2,
	"Diamonds": 1,
	"Hearts":   3,
	"Spades":   0,
}
