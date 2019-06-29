#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "RVExtension.h"

__attribute__((dllexport)) void RVExtension(char *output, int outputSize, const char *function) {
    goRVExtension(output, outputSize, function);
}