#include <stdio.h>
#include <math.h>

int main() {
    float distance;
    scanf("%f", &distance);

    int result = round((5280.0 / 4854.0) * 1000 * distance);
    printf("%d", result);
    return 0;
}