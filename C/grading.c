#include <stdio.h>

main() {
    int gradeLimits[5];
    char grades[5] = {'A', 'B', 'C', 'D', 'E'};

    for(int i = 0; i < 5; i++) {
        scanf("%d", &gradeLimits[i]);
    }

    int grade;
    scanf("%d", &grade);

    char result = 'F';
    for(int i = 0; i < 5; i++) {
        if (grade >= gradeLimits[i]) {
            result = grades[i];
            break;
        }
    }

    printf("%c", result);
}