#include <stdio.h>

long firstPositiveIntegerSum(int n) {
    long result = 0L;
    for (int i = 1; i <= n; i++) {
        result += i;
    }

    return result;
}

long firstOddOrEvenIntegerSum(int n, int isOdd) {
    long result = 0L;

    int aux = isOdd == 1 ? 1 : 2;
    for (int i = aux, totalNumbers = 0; totalNumbers < n; i += 2) {
        result += i;
        totalNumbers++;
    }

    return result;
}

int main() {
    int testCases;
    scanf("%d", &testCases);

    while(testCases--) {
        int dataSet, num;
        scanf("%d %d", &dataSet, &num);

        printf("%d %li %li %li\n", dataSet, firstPositiveIntegerSum(num), firstOddOrEvenIntegerSum(num, 1), firstOddOrEvenIntegerSum(num, 0));
    }
    return 0;
}