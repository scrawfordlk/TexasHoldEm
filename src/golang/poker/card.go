package poker

import (
	"fmt"
	"strings"
)

// Card represents a playing card with a suit and rank
type Card struct {
	Suit Suit
	Rank Rank
}

// NewCard creates a new Card from Suit and Rank
func NewCard(suit Suit, rank Rank) Card {
	return Card{Suit: suit, Rank: rank}
}

// NewCardFromString creates a new Card from a two-character string
func NewCardFromString(cardDesc string) (Card, error) {
	cardDesc = strings.ToUpper(cardDesc)
	if len(cardDesc) != 2 {
		return Card{}, fmt.Errorf("invalid card description: %s", cardDesc)
	}

	var suit Suit
	switch cardDesc[0] {
	case 'C':
		suit = CLUB
	case 'D':
		suit = DIAMOND
	case 'H':
		suit = HEART
	case 'S':
		suit = SPADE
	default:
		return Card{}, fmt.Errorf("invalid suit: %c", cardDesc[0])
	}

	var rank Rank
	switch cardDesc[1] {
	case '2':
		rank = TWO
	case '3':
		rank = THREE
	case '4':
		rank = FOUR
	case '5':
		rank = FIVE
	case '6':
		rank = SIX
	case '7':
		rank = SEVEN
	case '8':
		rank = EIGHT
	case '9':
		rank = NINE
	case 'T':
		rank = TEN
	case 'J':
		rank = JACK
	case 'Q':
		rank = QUEEN
	case 'K':
		rank = KING
	case 'A':
		rank = ACE
	default:
		return Card{}, fmt.Errorf("invalid rank: %c", cardDesc[1])
	}

	return Card{Suit: suit, Rank: rank}, nil
}

func (c Card) String() string {
	return fmt.Sprintf("%s-%s", c.Suit, c.Rank)
}
