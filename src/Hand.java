import java.util.Arrays;
import java.util.Comparator;

public class Hand {
    private Card[] cards;
    private HandVal handValue;

    public Hand(Card[] cards) {
        this.cards = cards;
        this.handValue = calcHandValue();
    }

    // constructor that takes a string describing 5 cards,
    // each with two characters, separated by at least one space
    // eg, "D2 HA ST CK S2"
    public Hand(String handDesc) {
        String[] cardDescs = handDesc.split("\\s+");
        if (cardDescs.length != 5) {
            throw new IllegalArgumentException("Invalid hand description: " + handDesc);
        }
        cards = new Card[5];
        for (int i = 0; i < 5; i++) {
            cards[i] = new Card(cardDescs[i]);
        }
        this.handValue = calcHandValue();
    }

    private HandVal calcHandValue() {
        // Placeholder for the actual hand value calculation logic
        // This will involve checking for various hand combinations
        // such as pairs, straights, flushes, etc.

        // Sort the cards by rank
        Arrays.sort(cards, Comparator.comparing(Card::getRank));

        // 1. Check for straight
        boolean isStraight = true;
        for (int i = 0; i < cards.length - 1; i++) {
            if (cards[i].getRank().getValue() != cards[i + 1].getRank().getValue() - 1) {
                isStraight = false;
                break;
            }
        }
        // Special case for Ace-low straight (A-2-3-4-5)
        if (!isStraight && cards[0].getRank() == Rank.TWO && cards[1].getRank() == Rank.THREE &&
                cards[2].getRank() == Rank.FOUR && cards[3].getRank() == Rank.FIVE && cards[4].getRank() == Rank.ACE) {
            isStraight = true;
        }

        // 2. checks for same suit
        boolean isSameSuit = Arrays.stream(cards).allMatch(card -> card.getSuit() == cards[0].getSuit());

        // 3. Straight flush or royal flush
        if (isStraight && isSameSuit) {
            // check royal flush
            if (cards[0].getRank() == Rank.TEN && cards[1].getRank() == Rank.JACK &&
                    cards[2].getRank() == Rank.QUEEN && cards[3].getRank() == Rank.KING &&
                    cards[4].getRank() == Rank.ACE) {
                return HandVal.ROYAL_FLUSH;
            } else {
                return HandVal.STRAIGHT_FLUSH;
            }
        }

        // 4. check for four of a kind
        for (int i = 0; i < cards.length - 3; i++) {
            int cardValue = cards[i].getRank().getValue();
            if (cardValue == cards[i + 1].getRank().getValue() &&
                    cardValue == cards[i + 2].getRank().getValue() &&
                    cardValue == cards[i + 3].getRank().getValue()) {
                return HandVal.POKER;
            }
        }

        // 5. check for full house

        // check triple
        boolean isThreeOfAKind = false;
        int isThreeOfAKindValue = -1;
        for (int i = 0; i < cards.length - 2; i++) {
            int cardValue = cards[i].getRank().getValue();
            if (cardValue == cards[i + 1].getRank().getValue() &&
                    cardValue == cards[i + 2].getRank().getValue()) {
                isThreeOfAKind = true;
                isThreeOfAKindValue = cardValue;
                break;
            }
        }

        // check pair
        boolean isPair = false;
        int isPairValue = -1;
        for (int i = 0; i < cards.length - 1; i++) {
            if (cards[i].getRank().getValue() == cards[i + 1].getRank().getValue()) {
                isPair = true;
                isPairValue = cards[i].getRank().getValue();
                if (isPairValue != isThreeOfAKindValue) {
                    break;
                }
            }
        }

        if (isThreeOfAKind && isPair && isThreeOfAKindValue != isPairValue) {
            return HandVal.FULL_HOUSE;
        }

        if (isSameSuit) {
            return HandVal.FLUSH;
        }

        if (isStraight) {
            return HandVal.STRAIGHT;
        }

        if (isThreeOfAKind) {
            return HandVal.THREE_OF_A_KIND;
        }

        // check for two pairs
        for (int i = 0; i < cards.length - 1; i++) {
            int cardValue = cards[i].getRank().getValue();
            if (cardValue == cards[i + 1].getRank().getValue()) {
                for (int j = i + 2; j < cards.length - 1; j++) {
                    int cardValue2 = cards[j].getRank().getValue();
                    if (cardValue2 == cards[j + 1].getRank().getValue()) {
                        return HandVal.TWO_PAIRS;
                    }
                }
            }
        }

        if (isPair) {
            return HandVal.PAIR;
        }

        // 5. Determine the highest value hand
        return HandVal.HIGH_CARD;
    }

    public Card[] getCards() {
        return cards;
    }

    public HandVal getHandValue() {
        return handValue;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        for (Card card : cards) {
            sb.append(card).append("\n"); // new line after each card
        }
        return sb.toString().trim();
    }
}
