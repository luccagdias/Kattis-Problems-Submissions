#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        int caloriesPerDay;
        scanf("%d", &caloriesPerDay);

        int result = caloriesPerDay / 400;
        printf("%d\n", caloriesPerDay % 400 == 0 ? result : result + 1);
    }

    return 0;
}