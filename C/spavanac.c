#include <stdio.h>

int main() {
    int hour, minute;
    scanf("%d %d", &hour, &minute);

    minute = minute - 45;
    if (minute < 0) {
        minute = 60 + minute;
        hour = hour == 0 ? 23 : --hour; 
    }

    printf("%d %d", hour, minute);
    return 0;
}