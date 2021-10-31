#include <stdio.h>
#include <math.h>

int main() {

    int a, b, c, d;
    scanf("%d %d %d %d", &a, &b, &c, &d);

    double semiPerimeter = (a + b + c + d) / 2.0;
    double result = sqrt((semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c) * (semiPerimeter - d));

    printf("%f", result);
    return 0;
}