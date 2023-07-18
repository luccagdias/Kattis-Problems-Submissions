import java.util.Scanner;

public class Metronome {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		int ticks = Integer.parseInt(scan.nextLine());

		double result = ticks / 4.0;
		System.out.println(String.format("%.2f", result));
	}
}
