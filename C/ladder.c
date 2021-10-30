#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int main() {
    int h, v;
    scanf("%d %d", &h, &v);

    float vInRad = v * (PI / 180);
    
    float ca = h / tan(vInRad);
    float result = ca / cos(vInRad);

    printf("%d", (int)result + 1);
    return 0;
}