#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);

    int nums[n];
    for (int i = 0; i < n; i++) {
        int num;
        scanf("%d", &num);

        nums[i] = num;
    }

    for (int i = n - 1; i >= 0; i--) {
        printf("%d\n", nums[i]);
    }


    return 0;
}