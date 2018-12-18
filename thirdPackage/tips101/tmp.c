#include <stdio.h>

int main() {
    struct test {
        int a;
        float b;
        char c;
        short d;
        int e[3];
        // double f;
        long g;
        short h[3];
    };
    struct test t;
    // printf() displays the string inside quotation
    printf("%d", sizeof(t));
    return 0;
}