#include <stdio.h>
#include <math.h>

int main() {
    int iterations;
    scanf("%d", &iterations);

    int result = 2;
    for (int i = 0; i < iterations; i++) {
        result += pow(2, i);
    }

    printf("%d", result * result);
    return 0;
}