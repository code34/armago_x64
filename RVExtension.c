#include <stdlib.h>

extern void goRVExtension(char *output, size_t outputSize, char *input);

__declspec(dllexport) void __stdcall _RVExtension(char *output, size_t outputSize, char *input) {
  goRVExtension(output, outputSize, input);
}

// do this for all the other exported functions