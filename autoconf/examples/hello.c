#include <stdio.h>
#include "config.h" // For PACKAGE and VERSION

int main() {
#ifdef ENABLE_GREETING
    printf("%s %s!\n", PACKAGE, VERSION);
#else
    printf("Hello, World!\n");
#endif
    return 0;
}
