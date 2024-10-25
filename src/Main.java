public class Main {
    public static void main(String[] args) {
        Card card1 = new Card(Suit.SPADE, Rank.ACE);
        Card card2 = new Card(Suit.HEART, Rank.TWO);
        Card card3 = new Card(Suit.DIAMOND, Rank.THREE);
        Card card4 = new Card(Suit.CLUB, Rank.FOUR);
        Card card5 = new Card(Suit.SPADE, Rank.FIVE);

        Hand hand = new Hand(new Card[] { card1, card2, card3, card4, card5 });
        System.out.println(hand);
        System.out.println("===============> " + hand.getHandValue());

        // royal flush example
        hand = new Hand("ST SJ SQ SK SA");
        System.out.println(hand);
        System.out.println("Royal flush");
        System.out.println("===============> " + hand.getHandValue());

        // straight flush example
        hand = new Hand("S6 S7 S8 S9 ST");
        System.out.println(hand);
        System.out.println("Straight flush");
        System.out.println("===============> " + hand.getHandValue());

        // four of a kind example
        hand = new Hand("S6 H6 D6 C6 ST");
        System.out.println(hand);
        System.out.println("four of a kind");
        System.out.println("===============> " + hand.getHandValue());

        // full house example
        hand = new Hand("S6 H6 D6 CT ST");
        System.out.println(hand);
        System.out.println("full house");
        System.out.println("===============> " + hand.getHandValue());

        // flush example
        hand = new Hand("S2 S4 S6 S8 ST");
        System.out.println(hand);
        System.out.println("flush");
        System.out.println("===============> " + hand.getHandValue());

        // straight example
        hand = new Hand("S6 H7 D8 C9 ST");
        System.out.println(hand);
        System.out.println("straight");
        System.out.println("===============> " + hand.getHandValue());

        // three of a kind example
        hand = new Hand("S6 H6 D6 C8 ST");
        System.out.println(hand);
        System.out.println("three of a kind");
        System.out.println("===============> " + hand.getHandValue());

        // two pairs example
        hand = new Hand("S6 H6 D8 C8 ST");
        System.out.println(hand);
        System.out.println("two pairs");
        System.out.println("===============> " + hand.getHandValue());

        // pair example
        hand = new Hand("S6 H7 D8 CT ST");
        System.out.println(hand);
        System.out.println("pair");
        System.out.println("===============> " + hand.getHandValue());

        // high card example
        hand = new Hand("S2 H3 D4 C6 S8");
        System.out.println(hand);
        System.out.println("high card");
        System.out.println("===============> " + hand.getHandValue());
    }
}
