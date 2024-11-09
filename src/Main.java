import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.FileNotFoundException;
import java.util.List;
import java.util.LinkedList;

public class Main {

    public static void main(String[] args) {
        String fileName = "Hands to be tested.txt";
        FileReader fReader = null;
        List<String> list = new LinkedList<>();
        try {
            try {
                fReader = new FileReader(fileName);
            } catch (FileNotFoundException e) {
                System.out.println(String.format("The file \"%s\" is not present in this directory", fileName));
                System.exit(1);
            }

            BufferedReader reader = new BufferedReader(fReader);
            String line = null;
            while ((line = reader.readLine()) != null) {
                line = stripFull(line);
                if (line.equals("")) {
                    continue;
                }

                list.add(line + "\n");
            }
            reader.close();
        } catch (IOException e) {
            e.printStackTrace();
        }

        HandVal[] values = new HandVal[] {
                HandVal.ROYAL_FLUSH,
                HandVal.STRAIGHT_FLUSH,
                HandVal.POKER,
                HandVal.FULL_HOUSE,
                HandVal.FLUSH,
                HandVal.STRAIGHT,
                HandVal.THREE_OF_A_KIND,
                HandVal.TWO_PAIRS,
                HandVal.PAIR,
                HandVal.HIGH_CARD
        };

        int i = -1;
        for (String line : list) {
            // test next HandVal
            if (line.charAt(0) == '/') {
                i++;
                continue;
            }

            testHand(line, values[i]);
        }
    }

    private static void testHand(String handDesc, HandVal expectedVal) {
        Hand hand = new Hand(handDesc);
        String result = expectedVal.toString();
        if (hand.getHandValue() == expectedVal) {
            result = result + " OK";
        } else {
            result = result + " NOT OK";
        }

        try {
            BufferedWriter writer = new BufferedWriter(new FileWriter("Output.txt", true));
            writer.write(result);
            writer.newLine();
            writer.close();
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(1);
        }

    }

    private static String stripFull(String str) {
        String stripped = str.strip();
        String newStr = "";
        for (char character : stripped.toCharArray()) {
            if (character == '"' || character == ',') {
                continue;
            }

            newStr = newStr + character;
        }

        return newStr;
    }
}
