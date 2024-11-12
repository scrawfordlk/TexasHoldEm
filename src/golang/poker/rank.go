package poker

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
