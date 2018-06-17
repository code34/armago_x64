package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export RVExtension
func RVExtension(output *C.char, outputsize C.ulonglong, input *C.char) {
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