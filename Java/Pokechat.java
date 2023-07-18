import java.util.Scanner;

public class Pokechat {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		String message = scan.nextLine();
		String IDs = scan.nextLine();
		
		StringBuffer sb = new StringBuffer();
		for (int i = 0; i < IDs.length(); i += 3) {
			String ID = IDs.substring(i, i + 3);
			int idx = Integer.parseInt(ID) - 1;

			sb.append(message.charAt(idx));
		}
		System.out.println(sb.toString());
	}
}
