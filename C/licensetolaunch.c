#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);

    long minAmout;
    scanf("%li", &minAmout);

    int result = 0;
    for(int i = 1; i < n; i++) {
        long amoutOnIDay;
        scanf("%li", &amoutOnIDay); 

        if (amoutOnIDay < minAmout) {
            minAmout = amoutOnIDay;
            result = i;
        }
    }

    printf("%li", result);
    return 0;
}