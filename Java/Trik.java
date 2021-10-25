import java.util.Scanner;

class Trik {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);

        String moves = scanner.nextLine();

        int[] result = new int[3];
        result[0] = 1;

        for(int i = 0; i < moves.length(); i++) {
            char move = moves.charAt(i);

            int position1 = 0;
            int position2 = 0;

            if (move == 'A') {
                position1 = 0;
                position2 = 1;
            }

            if (move == 'B') {
                position1 = 1;
                position2 = 2; 
            }

            if (move == 'C') {
                position1 = 0;
                position2 = 2;
            }

            int aux = result[position1];
            result[position1] = result[position2];
            result[position2] = aux; 

        }

        for(int i = 0; i < result.length; i++) {
            if (result[i] == 1) {
                System.out.println(i + 1);
                break;
            }
        }
    }
}