import java.util.Scanner;

public class AddingTrouble {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		int A = scan.nextInt();
		int B = scan.nextInt();
		int C = scan.nextInt();

		System.out.println(A + B == C ? "correct!" : "wrong!");
	}
}
