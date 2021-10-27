import java.util.Scanner;

class ShatteredCake {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        
        int W = scanner.nextInt();
        int N = scanner.nextInt();

        int totalArea = 0;
        for(int i = 0; i < N; i++) {
            int wi = scanner.nextInt();
            int li = scanner.nextInt();

            totalArea += wi * li;
        }

        System.out.println(totalArea / W);
    }
}