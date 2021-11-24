package models

import (
	"fmt"
	"testing"
)

func TestNewCard(t *testing.T) {
	// Arrange
	table := []struct {
		input    string
		expected *Card
	}{
		{"AH", &Card{Value: "Ace", Suit: "Hearts", Code: "AH"}},
		{"0K", nil},
		{"2P", nil},
	}

	for _, test := range table {
		testName := fmt.Sprintf("NewCard(%s)", test.input)
		t.Run(testName, func(t *testing.T) {
			// Arrange

			// Act
			card, err := NewCard(test.input)

			// Assert
			if test.expected != nil {
				if card.Value != test.expected.Value || card.Suit != test.expected.Suit || card.Code != test.expected.Code {
					t.Errorf("Expected %#v, got %#v", test.expected, card)
				}
			} else if err == nil {
				t.Errorf("Expected error for code %s", test.input)
			}
		})
	}
}

func TestGetPossibleCardCodes_WhenCalled_ReturnsFullDeck(t *testing.T) {
	// Arrange

	// Act
	codes := GetPossibleCardCodes()

	// Assert
	if len(codes) != 52 {
		t.Errorf("Expected 52 card codes, got %d", len(codes))
	}
}
