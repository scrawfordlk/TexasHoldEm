public class Main {
    public static void main(String[] args) {
        Card card1 = new Card(Rank.ACE, Suit.SPADE);
        Card card2 = new Card(Rank.TWO, Suit.HEART);
        Card card3 = new Card(Rank.THREE, Suit.DIAMOND);
        Card card4 = new Card(Rank.FOUR, Suit.CLUB);
        Card card5 = new Card(Rank.FIVE, Suit.SPADE);

        Hand hand = new Hand(new Card[] {card1, card2, card3, card4, card5});

        System.out.println(hand);
    }
}