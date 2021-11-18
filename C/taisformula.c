#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);

    long long lastTime = 0LL;
    float lastValue;
    scanf("%lld %f", &lastTime, &lastValue);

    double result;
    for (int i = 0; i < n - 1; i++) {
        long long time = 0LL;
        float value;

        scanf("%lld %f", &time, &value);

        result += (((lastValue + value) / 2) * (time - lastTime)) / 1000;

        lastTime = time;
        lastValue = value;
    }

    printf("%lf", result);
    return 0;
}