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

        // Example logic (not complete):
        // 1. Sort the cards by rank
        Arrays.sort(cards, Comparator.comparing(Card::getRank));

        // check for one pair
        boolean isOnePair = false;
        for (int i = 0; i < cards.length - 1; i++) {
            if (cards[i].getRank().getValue() == cards[i + 1].getRank().getValue()) {
                isOnePair = true;
                break;
            }
        }
        if (isOnePair) {
            return HandVal.PAIR;
        }

        // 2. Check for flush
        boolean isFlush = Arrays.stream(cards).allMatch(card -> card.getSuit() == cards[0].getSuit());
        if (isFlush) {
            return HandVal.FLUSH;
        }
        // 3. Check for straight
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
        if (isStraight) {
            return HandVal.STRAIGHT;
        }
        // check for poker
        boolean isPoker = false;
        for (int i = 0; i < cards.length - 3; i++) {
            if (cards[i].getRank().getValue() == cards[i + 1].getRank().getValue() &&
                    cards[i].getRank().getValue() == cards[i + 2].getRank().getValue() &&
                    cards[i].getRank().getValue() == cards[i + 3].getRank().getValue()) {
                isPoker = true;
                break;
            }
        }
        if (isPoker) {
            return HandVal.POKER;
        }
        // 4. Check for pairs, three of a kind, four of a kind

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
            sb.append(card).append("\n");  // new line after each card
        }
        return sb.toString().trim();
    }
}
