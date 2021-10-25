import java.util.Scanner;

class TransitWoes {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        
        int s = 0, t = 0, n = 0;
        s = scanner.nextInt();
        t = scanner.nextInt();
        n = scanner.nextInt();

        int[] walkTimes = new int[n + 1];
        for(int i = 0; i <= n; i++) {
            walkTimes[i] += scanner.nextInt();
        }

        int[] busRideTimes = new int[n];
        for(int i = 0; i < n; i++) {
            busRideTimes[i] += scanner.nextInt();
        }

        int[] busIntervals = new int[n];
        for(int i = 0; i < n; i++) {
            busIntervals[i] += scanner.nextInt();
        }

        int totalTime = 0;
        for(int i = 0; i < n; i++) {
            totalTime = walkTimes[i];

            if (totalTime <= busIntervals[i]) {
                totalTime += busIntervals[i] - totalTime;
            } else {
                totalTime += 3 - (totalTime % busIntervals[i]);
            }

            totalTime += busRideTimes[i];
        }

        totalTime += walkTimes[n];

        String result = totalTime <= (t - s) ? "yes" : "no";
        System.out.println(result);
    }
}