package poker

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
