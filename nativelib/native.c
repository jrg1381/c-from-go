#include <stdio.h>
#include <malloc.h>
#include "native.h"

#define BUFLEN 32

char* some_function(int param1, char* param2)
{
    char *buf = malloc(BUFLEN * sizeof(char));
    snprintf(buf, BUFLEN, "%d-%s", param1, param2);
    return buf;
}