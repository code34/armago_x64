package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
*/
import "C"

import "unsafe"
import "fmt"

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