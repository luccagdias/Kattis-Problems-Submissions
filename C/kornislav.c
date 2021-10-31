#include <stdio.h>

void sortArray(int array[]) {
    int lastUnsortedPosition = 4;

    while (lastUnsortedPosition > 0) {
        for (int i = 0; i < lastUnsortedPosition - 1; i++) {
            if (array[i] > array[i + 1]) {
                int aux = array[i];
                array[i] = array[i + 1];
                array[i + 1] = aux;
            }
        }

        lastUnsortedPosition--;
    }
}

int main() {

    int nums[4];
    int a, b, c, d;
    scanf("%d %d %d %d", &nums[0], &nums[1], &nums[2], &nums[3]);

    sortArray(nums);
    printf("%d", nums[0] * nums[2]);
    return 0;
}