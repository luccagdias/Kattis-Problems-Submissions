#include <stdio.h>

int main() {
    int n, sideLength = 1, height = 0;
    scanf("%d", &n);
 
    while(n >= sideLength * sideLength) {
        n -= sideLength * sideLength;
        sideLength += 2;
        height++;
    }

    printf("%d", height);
    return 0;
}