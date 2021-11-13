import java.util.Scanner;
import java.lang.Math;
import java.lang.StringBuilder;

class Encodedmessage {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        
        int n = scanner.nextInt();
        scanner.nextLine();

        while(n-- > 0) {
            StringBuilder result = new StringBuilder();
            String input = scanner.nextLine();
            int iterationsToMake = (int) Math.sqrt((double) input.length());

            for (int i = iterationsToMake; i > 0; i--) {
                int position = i - 1;
                for (int j = iterationsToMake; j > 0; j--) {
                    result.append(input.charAt(position));
                    position += iterationsToMake;
                }
            }

            System.out.println(result);
        }
    }
}