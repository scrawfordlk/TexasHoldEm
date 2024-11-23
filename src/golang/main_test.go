package main

import (
	"golang/poker"
	"testing"
)

// TestHands validates poker hand evaluation.
func TestHands(t *testing.T) {
	tests := []struct {
		handDesc       string
		expectedHandVal poker.HandValue
	}{
		// Royal Flush
		{"CT CJ CQ CK CA", poker.ROYAL_FLUSH},

		// Straight Flush
		{"D8 DQ DJ DT D9", poker.STRAIGHT_FLUSH},
		{"H8 HT HJ H7 H9", poker.STRAIGHT_FLUSH},

		// Four of a Kind
		{"HT SQ ST DT CT", poker.FOUR_OF_A_KIND},
		{"HT SK ST DT CT", poker.FOUR_OF_A_KIND},
		{"H8 SQ S8 D8 C8", poker.FOUR_OF_A_KIND},
		{"H7 SK S7 D7 C7", poker.FOUR_OF_A_KIND},

		// Full House
		{"H2 SQ C2 D2 CQ", poker.FULL_HOUSE},
		{"H2 SJ C2 D2 CJ", poker.FULL_HOUSE},

		// Flush
		{"HK HQ H2 H4 H5", poker.FLUSH},
		{"D5 D4 D2 DQ DK", poker.FLUSH},

		// Straight
		{"H3 S7 H5 D6 H4", poker.STRAIGHT},
		{"C9 CT SJ D7 H8", poker.STRAIGHT},
		{"H4 S5 HA D3 H2", poker.STRAIGHT},

		// Three of a Kind
		{"H2 SQ S2 D2 CK", poker.THREE_OF_A_KIND},
		{"H2 S7 S2 D2 C9", poker.THREE_OF_A_KIND},
		{"H2 S8 S2 D2 C9", poker.THREE_OF_A_KIND},

		// Two Pairs
		{"H5 SQ C5 DT CT", poker.TWO_PAIR},
		{"H9 SQ C9 DT CT", poker.TWO_PAIR},

		// One Pair
		{"H3 S8 H5 D8 CA", poker.PAIR},
		{"S4 DA H3 CA HT", poker.PAIR},

		// High Card
		{"H3 S8 H5 DK CA", poker.HIGH_CARD},
		{"H3 S8 H5 DK CT", poker.HIGH_CARD},
		{"H3 S8 H5 DK C2", poker.HIGH_CARD},
	}

	for _, tt := range tests {
		t.Run(tt.handDesc, func(t *testing.T) {
			hand, err := poker.NewHandFromString(tt.handDesc)
			if err != nil {
				t.Fatalf("Error creating hand from %q: %v", tt.handDesc, err)
			}

			if hand.HandValue != tt.expectedHandVal {
				t.Errorf("Hand %q: got %v, want %v", tt.handDesc, hand.HandValue, tt.expectedHandVal)
			}
		})
	}
}
