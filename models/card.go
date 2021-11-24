package models

import (
	"errors"
	"strings"
)

func NewCard(code string) (*Card, error) {
	rank, suit, err := decodeCardCode(code)
	if err != nil {
		return nil, err
	}

	return &Card{
		Value: rank,
		Suit:  suit,
		Code:  code,
	}, nil
}

func GetPossibleCardCodes() []string {
	var result = make([]string, 0, len(ranks)*len(suits))
	for suit := range suits {
		for rank := range ranks {
			result = append(result, rank+suit)
		}
	}

	return result
}

type Card struct {
	Value string `json:"value" example:"Ace"`
	Suit  string `json:"suit" example:"Hearts"`
	Code  string `json:"code" example:"AH"`
}

func decodeCardCode(code string) (rank string, suit string, err error) {
	suitCode := code[len(code)-1:]
	suit, exists := suits[suitCode]
	if !exists {
		return "", "", errors.New("No suit exists with code " + suitCode)
	}

	rankCode := strings.TrimSuffix(code, suitCode)
	rank, exists = ranks[rankCode]
	if !exists {
		return "", "", errors.New("No rank exists with code " + rankCode)
	}

	return
}

var ranks = map[string]string{
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"J":  "Jack",
	"Q":  "Queen",
	"K":  "King",
	"A":  "Ace",
}

var suits = map[string]string{
	"C": "Clubs",
	"D": "Diamonds",
	"H": "Hearts",
	"S": "Spades",
}
