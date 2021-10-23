#include <stdio.h>

main() {
    int n, input, result;
    scanf("%d", &n);

    result = 0;
    for(int i = 0; i < n; i++) {
        scanf("%d", &input);
        result += input;
    }

    for(int i = 0; i < n - 1; i++) {
        scanf("%d", &input);
        result -= input;
    }

    printf("%d", result);
}