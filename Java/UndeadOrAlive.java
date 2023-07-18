import java.util.Scanner;

public class UndeadOrAlive {
	static String resolve(boolean smileyFace, boolean frownyFace) {
		if (smileyFace && frownyFace) 
			return "double agent";

		if (!smileyFace && !frownyFace)
			return "machine";

		if (smileyFace)
			return "alive";

		return "undead";
	}

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		String message = scan.nextLine();

		boolean smileyFace = message.contains(":)");
		boolean frownyFace = message.contains(":(");
		
		System.out.println(resolve(smileyFace, frownyFace));
	}
}
