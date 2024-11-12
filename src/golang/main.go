package main

import (
	"fmt"
	"golang/poker"
)

func main() {
	// Test Royal Flush
	testHand("CT CJ CQ CK CA", poker.ROYAL_FLUSH)

	// Test Straight Flush
	testHand("D8 DQ DJ DT D9", poker.STRAIGHT_FLUSH)
	testHand("H8 HT HJ H7 H9", poker.STRAIGHT_FLUSH)

	// Test Four of a Kind (Poker)
	testHand("HT SQ ST DT CT", poker.FOUR_OF_A_KIND)
	testHand("HT SK ST DT CT", poker.FOUR_OF_A_KIND)
	testHand("H8 SQ S8 D8 C8", poker.FOUR_OF_A_KIND)
	testHand("H7 SK S7 D7 C7", poker.FOUR_OF_A_KIND)

	// Test Full House
	testHand("H2 SQ C2 D2 CQ", poker.FULL_HOUSE)
	testHand("D2 DJ D2 D2 DJ", poker.FULL_HOUSE)
	testHand("H2 SJ C2 D2 CJ", poker.FULL_HOUSE)
	testHand("D2 DJ D2 D2 DJ", poker.FULL_HOUSE)

	// Test Flush
	testHand("HK HQ H2 H4 H5", poker.FLUSH)
	testHand("HK HQ H2 H4 H5", poker.FLUSH)
	testHand("D5 D4 D2 DQ DK", poker.FLUSH)

	// Test Straight
	testHand("H3 S7 H5 D6 H4", poker.STRAIGHT)
	testHand("C9 CT SJ D7 H8", poker.STRAIGHT)
	testHand("H4 S5 HA D3 H2", poker.STRAIGHT)

	// Test Three of a Kind
	testHand("H2 SQ S2 D2 CK", poker.THREE_OF_A_KIND)
	testHand("H2 S7 S2 D2 C9", poker.THREE_OF_A_KIND)
	testHand("H2 S8 S2 D2 C9", poker.THREE_OF_A_KIND)

	// Test Two Pairs
	testHand("H5 SQ C5 DT CT", poker.TWO_PAIR)
	testHand("D5 DK S5 DT DT", poker.TWO_PAIR)
	testHand("H9 SQ C9 DT CT", poker.TWO_PAIR)
	testHand("D8 DK S8 DT DT", poker.TWO_PAIR)

	// Test One Pair
	testHand("H3 S8 H5 D8 CA", poker.PAIR)
	testHand("S4 DA H3 CA HT", poker.PAIR)

	// Test High Card
	testHand("H3 S8 H5 DK CA", poker.HIGH_CARD)
	testHand("H3 S8 H5 DK CT", poker.HIGH_CARD)
	testHand("H3 S8 H5 DK C2", poker.HIGH_CARD)
}

// testHand is a helper function that tests a hand and prints if it's correct
func testHand(handDesc string, expectedHandVal poker.HandValue) {
	hand, err := poker.NewHandFromString(handDesc)
	if err != nil {
		fmt.Printf("Error creating hand: %v\n", err)
		return
	}
	fmt.Println(hand)
	fmt.Printf("  =>   %s\n", hand.HandValue)

	if hand.HandValue == expectedHandVal {
		fmt.Println("  OK")
	} else {
		fmt.Println("  NOT OK")
	}
	fmt.Println("___________________________")
}
