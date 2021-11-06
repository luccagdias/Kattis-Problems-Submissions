import java.util.Scanner;
import java.lang.StringBuilder;

class Avion {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        StringBuilder result = new StringBuilder();

        for (int i = 0; i < 5; i++) {
            String blimp = scanner.nextLine();

            CharSequence fbi = "FBI";
            if (blimp.contains(fbi)) {
                result.append(i + 1 + " ");
            }
        }

        System.out.println(result.length() == 0 ? "HE GOT AWAY!" : result);
    }
}