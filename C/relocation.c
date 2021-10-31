#include <stdio.h>
#include <stdlib.h>

int main() {
    int n, q;
    scanf("%d %d", &n, &q);

    int nums[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &nums[i]);
    }

    for (int i = 0; i < q; i++) {
        int a, b, c;
        scanf("%d %d %d", &a, &b, &c);

        if (a == 1) {
            nums[b - 1] = c;
        } else {
            printf("%d\n", abs(nums[b - 1] - nums[c - 1]));
        }
    }

    return 0; 
}