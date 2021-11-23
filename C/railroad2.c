#include <stdio.h>

int main() {
    int x, y;
    scanf("%d %d", &x, &y);

    printf("%s", y % 2 != 0 ? "impossible" : "possible");
    return 0;
}