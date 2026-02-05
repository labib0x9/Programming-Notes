#include<stdio.h>
#include<stdbool.h>

struct A {
    bool a;
    int b;
    char c;
    long long d;
    char e;
};


struct B {
    bool a;
    int b;
    char c;
    long long d;
    char e;
}__attribute__((packed));


int main() {

    printf("%zu\n", sizeof(struct A));
    printf("%zu\n", sizeof(struct B));

    printf("%zu\n", __offsetof(struct A, b));
    printf("%zu\n", __offsetof(struct B, b));

    return 0;
}