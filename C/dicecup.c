#include <stdio.h>

int main() {    
    int n, m;
    scanf("%d %d", &n, &m);

    int min = n < m ? n : m;
    int max = n > m ? n : m;

    for (int i = min + 1; i <= max + 1; i++) {
        printf("%d\n", i);
    }

    return 0;
}