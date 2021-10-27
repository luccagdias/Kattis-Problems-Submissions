#include <stdio.h>

main() {
    int a, b;

    scanf("%d", &a);
    scanf("%d", &b);

    if (a > b) {
        printf("%d %d", b, a);
    } else {
        printf("%d %d", a, b);
    }
}