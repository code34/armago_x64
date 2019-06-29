package main


/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lRVExtension

__attribute__((dllexport)) void RVExtension(char *output, int outputSize, const char *function);

*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export goRVExtension
func goRVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	temp := fmt.Sprintf("Hello %s!", C.GoString(input))
	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func main() {}