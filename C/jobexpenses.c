#include <stdio.h>
#include <stdlib.h>

int main(){
    int n;
    scanf("%d", &n);

    int total = 0;
    for (int i = 0; i < n; i++) {
        int num;
        scanf("%d", &num);

        if (num < 0) {
            total += num;
        }
    }

    printf("%d", abs(total));
    return 0;
}