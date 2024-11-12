package poker

import (
	"errors"
	"fmt"
	//"sort"
	"strings"
)

// Hand represents a poker hand of five cards
type Hand struct {
	Cards     []Card
	HandValue HandValue
}

// NewHand creates a Hand from an array of Card
func NewHand(cards []Card) (Hand, error) {
	if len(cards) != 5 {
		return Hand{}, errors.New("a hand must contain exactly 5 cards")
	}
	hand := Hand{Cards: cards}
	hand.HandValue = hand.calcHandValue()
	return hand, nil
}

// NewHandFromString creates a Hand from a string describing 5 cards, e.g., "D2 HA ST CK S2"
func NewHandFromString(handDesc string) (Hand, error) {
	cardDescs := strings.Fields(handDesc)
	if len(cardDescs) != 5 {
		return Hand{}, fmt.Errorf("invalid hand description: %s", handDesc)
	}
	cards := make([]Card, 5)
	for i, cardDesc := range cardDescs {
		card, err := NewCardFromString(cardDesc)
		if err != nil {
			return Hand{}, err
		}
		cards[i] = card
	}
	return NewHand(cards)
}

// calcHandValue determines the hand's rank by checking for pairs, straights, flushes, etc.
func (h *Hand) calcHandValue() HandValue {
	// Sorting and evaluation logic, as in the full code example
	// ...
	return HIGH_CARD // Placeholder for actual logic
}

func (h Hand) String() string {
	var sb strings.Builder
	for _, card := range h.Cards {
		sb.WriteString(fmt.Sprintf("%v\n", card))
	}
	return strings.TrimSpace(sb.String())
}
