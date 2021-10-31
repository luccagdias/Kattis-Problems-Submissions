import java.util.Scanner;
import java.util.HashSet;
import java.util.Set;

class Modulo {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        Set<Integer> modulos = new HashSet<>();

        for (int i = 0; i < 10; i++) {
            int n = scanner.nextInt();

            modulos.add(n % 42);
        }

        System.out.println(modulos.size());
    }
}