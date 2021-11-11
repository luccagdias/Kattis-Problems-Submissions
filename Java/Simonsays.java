import java.util.Scanner;

class Simonsays {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        scanner.nextLine();

        for (int i = 0; i < n; i++) {
            String line = scanner.nextLine();

            if (line.contains("Simon says")) {
                System.out.println(line.replace("Simon says", ""));
            }
        }
    }
}