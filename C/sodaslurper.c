#include <stdio.h>

int main() {
    int bottles, bottlesFound, cost;
    scanf("%d %d %d", &bottles, &bottlesFound, &cost);

    int result = 0, totalBottles = bottles + bottlesFound;
    while (totalBottles >= cost) {
        result += totalBottles / cost;
        totalBottles = (totalBottles / cost) + (totalBottles % cost);
    }

    printf("%d", result);
    return 0;
}