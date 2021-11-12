import java.util.Scanner;

class Smil {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        String symbols = scanner.nextLine();

        int size = symbols.length();
        for (int i = 0; i < size - 1; i++) {
            if (symbols.charAt(i) == ':' || symbols.charAt(i) == ';') {
                if (symbols.charAt(i + 1) == ')') {
                    System.out.println(i);
                }
            }
        }

        for (int i = 0; i < size - 2; i++) {
            if (symbols.charAt(i) == ':' || symbols.charAt(i) == ';') {
                if (symbols.charAt(i + 1) == '-') {
                    if (symbols.charAt(i + 2) == ')') {
                        System.out.println(i);
                    }
                }
            }
        }
    }
}