#include <stdio.h>

int main() {
    int left, right;
    scanf("%d %d", &left, &right);

    if (left == 0 && right == 0) {
        printf("Not a moose");    
    } else {
        if (left == right) {
            printf("Even %d", left + right);
        } else {
            int result = left > right ? left : right;
            printf("Odd %d", 2 * result);
        }
    }
    
    return 0;
}