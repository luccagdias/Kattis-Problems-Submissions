#include <stdio.h>

main() {
    int h, b;
    scanf("%d %d", &h, &b);

    float result = (h * b) / (float)2; 
    printf("%.1f", result);
}