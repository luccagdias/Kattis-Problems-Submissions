#include <stdio.h>
#include <stdbool.h>

int main() {

    char string[30];
    scanf("%s", string);

    bool hiss = false;
    for (int i = 0; i < 29; i++) {
        if (string[i] == 's' && string[i + 1] == 's') {
            hiss = true;
            break;
        }
    }

    printf(hiss ? "hiss" : "no hiss");
    return 0;
}