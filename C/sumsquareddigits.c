#include <stdio.h>

long sumSquaredDigits(int base, long num) {
    long remainder, result = 0, quocient = num;

    while(quocient > 0) {
        remainder = quocient % base;
        quocient = quocient / base;

        result += remainder * remainder;
    }

    return result;
}

int main() {
    int n;
    scanf("%d", &n);

    while(n--) {
        long num;
        int dataSetNumber, base;
        scanf("%d %d %li", &dataSetNumber, &base, &num);

        printf("%d %li\n", dataSetNumber, sumSquaredDigits(base, num));
    }

    return 0;
}