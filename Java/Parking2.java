import java.util.Scanner;
import java.util.Arrays;

class Parking2 {

    private static int[] createArrayWithInputs(Scanner scanner, int n) {
        int[] nums = new int[n];
        for (int j = 0; j < n; j++) {
            nums[j] = scanner.nextInt();
        }

        Arrays.sort(nums);
        return nums;
    }

    private static int solve(int[] nums) {
        int result = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            result += nums[i + 1] - nums[i];
        }

        result += nums[nums.length - 1] - nums[0];

        return result;
    }

    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        
        int testCases = scanner.nextInt();
        for (int i = 0; i < testCases; i++) {
            int n = scanner.nextInt();

            int[] nums = createArrayWithInputs(scanner, n);

            System.out.println(solve(nums));
        }
    }
}