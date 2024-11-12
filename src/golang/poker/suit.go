package poker

//import "fmt"

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
