#include <stdio.h>

void sortArray(int array[], int size) {
    int lastUnsortedPosition = size - 1;

    while (lastUnsortedPosition > 0) {
        for (int i = 0; i < lastUnsortedPosition; i++) {
            if (array[i] > array[i + 1]) {
                int aux = array[i];
                array[i] = array[i + 1];
                array[i + 1] = aux;
            }
        }

        lastUnsortedPosition--;
    }
}

int findNumPosition(int array[], int arraySize, int num) {
    for (int i = 0; i < arraySize; i++) {
        if (array[i] == num) {
            return i;
        }
    }
}

int main() {
    int testCases;
    scanf("%d", &testCases);

    while (testCases != 0) {
        int size = testCases;

        int firstList[size];
        int reference[size];
        for (int i = 0; i < size; i++) {
            scanf("%d", &firstList[i]);
            reference[i] = firstList[i];
        }

        int secondList[size];
        for (int i = 0; i < size; i++) {
            scanf("%d", &secondList[i]);
        }

        sortArray(firstList, size);
        sortArray(secondList, size);

        int result[size];
        for (int i = 0; i < size; i++) {
            int numToFind = reference[i];
            int position = findNumPosition(firstList, size, numToFind);

            result[i] = secondList[position];
        }

        for (int i = 0; i < size; i++) {
            printf("%d\n", result[i]);
        }

        printf("\n");
        scanf("%d", &testCases);
    }

    return 0;
}