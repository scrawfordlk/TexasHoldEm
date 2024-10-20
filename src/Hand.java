public class Hand {
    private Card[] cards;

    public Hand(Card[] cards) {
        this.cards = cards;
    }

    public Card[] getCards() {
        return cards;
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
