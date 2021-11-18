#include <stdio.h>
#include <math.h>

int main() {
    float input;
    scanf("%f", &input);
    
    printf("%.6f", pow(input, 1 / input));
    return 0;
}