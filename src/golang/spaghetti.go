package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Suit represents the suit of a card
type Suit int

const (
	CLUB Suit = iota
	DIAMOND
	HEART
	SPADE
)

func (s Suit) String() string {
	return [...]string{"CLUB", "DIAMOND", "HEART", "SPADE"}[s]
}

// Rank represents the rank of a card
type Rank int

const (
	TWO Rank = iota + 2 // TWO = 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r-2]
}

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

// HandValue represents the rank of a poker hand
type HandValue int

const (
	HIGH_CARD HandValue = iota
	PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
	ROYAL_FLUSH
	NOT_VALID
)

func (hv HandValue) String() string {
	return [...]string{
		"High Card", "Pair", "Two Pair", "Three of a Kind", "Straight", "Flush",
		"Full House", "Four of a Kind", "Straight Flush", "Royal Flush", "Not Valid",
	}[hv]
}

// Value returns the integer value associated with a HandValue
func (hv HandValue) Value() int {
	return int(hv)
}

// Hand represents a poker hand of five cards
type Hand struct {
	Cards    []Card
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
	// Sort cards by rank
	sort.Slice(h.Cards, func(i, j int) bool {
		return h.Cards[i].Rank < h.Cards[j].Rank
	})

	// Check for straight
	isStraight := true
	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].Rank != h.Cards[i+1].Rank-1 {
			isStraight = false
			break
		}
	}
	// Special case for Ace-low straight (A-2-3-4-5)
	if !isStraight && h.Cards[0].Rank == TWO && h.Cards[1].Rank == THREE &&
		h.Cards[2].Rank == FOUR && h.Cards[3].Rank == FIVE && h.Cards[4].Rank == ACE {
		isStraight = true
	}

	// Check for same suit (flush)
	isFlush := true
	for _, card := range h.Cards {
		if card.Suit != h.Cards[0].Suit {
			isFlush = false
			break
		}
	}

	// Straight flush or royal flush
	if isStraight && isFlush {
		if h.Cards[0].Rank == TEN && h.Cards[1].Rank == JACK &&
			h.Cards[2].Rank == QUEEN && h.Cards[3].Rank == KING &&
			h.Cards[4].Rank == ACE {
			return ROYAL_FLUSH
		}
		return STRAIGHT_FLUSH
	}

	// Check for four of a kind
	for i := 0; i < len(h.Cards)-3; i++ {
		if h.Cards[i].Rank == h.Cards[i+1].Rank &&
			h.Cards[i].Rank == h.Cards[i+2].Rank &&
			h.Cards[i].Rank == h.Cards[i+3].Rank {
			return FOUR_OF_A_KIND
		}
	}

	// Check for full house
	isThreeOfAKind := false
	threeOfAKindRank := -1
	isPair := false
	pairRank := -1

	for i := 0; i < len(h.Cards)-2; i++ {
		if h.Cards[i].Rank == h.Cards[i+1].Rank &&
			h.Cards[i].Rank == h.Cards[i+2].Rank {
			isThreeOfAKind = true
			threeOfAKindRank = int(h.Cards[i].Rank)
			break
		}
	}

	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].Rank == h.Cards[i+1].Rank {
			isPair = true
			pairRank = int(h.Cards[i].Rank)
			if pairRank != threeOfAKindRank {
				break
			}
		}
	}

	if isThreeOfAKind && isPair && threeOfAKindRank != pairRank {
		return FULL_HOUSE
	}

	if isFlush {
		return FLUSH
	}

	if isStraight {
		return STRAIGHT
	}

	if isThreeOfAKind {
		return THREE_OF_A_KIND
	}

	// Check for two pairs
	pairCount := 0
	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].Rank == h.Cards[i+1].Rank {
			pairCount++
			i++ // skip next card as it's part of this pair
		}
	}
	if pairCount == 2 {
		return TWO_PAIR
	}

	if isPair {
		return PAIR
	}

	// Otherwise, it is a high card
	return HIGH_CARD
}

// String returns a string representation of the Hand
func (h Hand) String() string {
	var sb strings.Builder
	for _, card := range h.Cards {
		sb.WriteString(fmt.Sprintf("%v\n", card))
	}
	return strings.TrimSpace(sb.String())
}

func main() {
	//
