package poker

import (
	"errors"
	"fmt"
	"sort"
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

func (h Hand) String() string {
	var sb strings.Builder
	for _, card := range h.Cards {
		sb.WriteString(fmt.Sprintf("%v\n", card))
	}
	return strings.TrimSpace(sb.String())
}
