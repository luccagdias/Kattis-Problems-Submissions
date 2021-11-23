#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);

    while(n != -1) {
        int distance, initialTime = 0, finalTime, result = 0;
        
        for (int i = 0; i < n; i++) {
            scanf("%d %d", &distance, &finalTime);

            result += distance * (finalTime - initialTime);
            initialTime = finalTime;
        }

        printf("%d miles\n", result);
        scanf("%d", &n);
    }

    return 0;
}