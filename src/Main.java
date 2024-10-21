public class Main {
    public static void main(String[] args) {
        Card card1 = new Card(Suit.SPADE, Rank.ACE);
        Card card2 = new Card(Suit.HEART, Rank.TWO);
        Card card3 = new Card(Suit.DIAMOND, Rank.THREE);
        Card card4 = new Card(Suit.CLUB, Rank.FOUR);
        Card card5 = new Card(Suit.SPADE, Rank.FIVE);

        Hand hand = new Hand(new Card[] {card1, card2, card3, card4, card5});
        System.out.println(hand);
        System.out.println("===============>   " + hand.getHandValue());

        hand= new Hand("S6    H2 D3  C4     S5");
        System.out.println(hand);
        System.out.println("===============>   " + hand.getHandValue());

        hand= new Hand("SK S5 HK DK CK");
        System.out.println(hand);
        System.out.println("===============>   " + hand.getHandValue());

        hand= new Hand("S4 S8 S2 SK S3");
        System.out.println(hand);
        System.out.println("===============>   " + hand.getHandValue());

        hand= new Hand("S6 H3 CA H8 C6");
        System.out.println(hand);
        System.out.println("===============>   " + hand.getHandValue());
    }
}