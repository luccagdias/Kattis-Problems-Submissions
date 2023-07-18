import java.util.Scanner;

public class Betting {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		double a = scan.nextInt();

		double opt1 = 1 / (a / 100.0);
		double opt2 = 1 / ((100 - a) / 100.0);
		System.out.println(String.format("%.10f\n%.10f", opt1, opt2));
	}
}
