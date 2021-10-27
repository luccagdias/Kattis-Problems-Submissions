import java.util.Scanner;

class Speeding {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        
        int n = scanner.nextInt();

        int lastTime = 0, lastDistance = 0, max = 0;
        for (int i = 0; i < n; i++) {
            int time = scanner.nextInt();
            int distance = scanner.nextInt();

            int t = lastTime - time;
            int d = lastDistance - distance;

            lastTime = time;
            lastDistance = distance;

            if (t != 0 && (d / t) > max) {
                max = d / t;
            }            
        }

        System.out.println(max);
    }
}