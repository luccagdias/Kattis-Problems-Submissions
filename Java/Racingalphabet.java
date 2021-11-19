import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.lang.Character;
import java.lang.Math;

class Racingalphabet {
    final static Double SIZE_OF_EACH_POINT =  6.731984257692413;
    final static List<Character> ALPHABET = new ArrayList<>(Arrays.asList('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', ' ', '\''));

    public static int findShortestDistance(Character char1, Character char2) {
        int fromIndex = 0;
        int toIndex = 0;

        int char1Index = ALPHABET.indexOf(char1);
        int char2Index = ALPHABET.indexOf(char2);

        if (char1Index < char2Index) {
            fromIndex = char1Index;
            toIndex = char2Index;
        } else {
            fromIndex = char2Index;
            toIndex = char1Index;
        }

        int result = Math.min((toIndex - fromIndex), Math.abs((28 - toIndex) + fromIndex));

        return result;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        scanner.nextLine();

        while(n-- > 0) {
            String line = scanner.nextLine();

            Double result = 1.0;
            for (int i = 0; i < line.length() - 1; i++) {
                int distance = findShortestDistance(line.charAt(i), line.charAt(i + 1));

                result += ((distance * SIZE_OF_EACH_POINT) / 15) + 1;
            }

            System.out.println(result);
        }
    }
}