#include <stdio.h>

main() {
    int n, num, total;
    scanf("%d", &n);

    total = 0;
    for (int i = 0; i < n; i++) {
        scanf("%d", &num);
        total += num;
    }

    printf("%d\n", total);
}