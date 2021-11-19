#include <stdio.h>

int main() {
    int testCases;
    scanf("%d", &testCases);

    while(testCases--) {
        int initialNumOfTurtles = 0, finalNumOfTurtles = -1, result = 0;
        scanf("%d %d", &initialNumOfTurtles, &finalNumOfTurtles);

        while (finalNumOfTurtles != 0) {
            if (finalNumOfTurtles > initialNumOfTurtles * 2) {
                result += finalNumOfTurtles - (initialNumOfTurtles * 2);
            }

            initialNumOfTurtles = finalNumOfTurtles;
            scanf("%d", &finalNumOfTurtles);
        }

        printf("%d\n", result);
    }

    return 0;
}