import java.util.Scanner;

public class GCVWR {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		int G = scan.nextInt();
		int T = scan.nextInt();
		int N = scan.nextInt();

		int items = 0;
		for(int i = 0; i < N; i++) {
			items += scan.nextInt();
		}
		
		int max = (int) ((G - T) * 0.9);
		System.out.println(max - items);
	}
}
