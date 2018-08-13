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

//export RVExtensionVersion
func RVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("Version 1.0")
	defer C.free(unsafe.Pointer(result))
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), outputsize - 1)
}

//export RVExtensionArgs
func RVExtensionArgs(output *C.char, outputsize C.size_t, input *C.char, argv **C.char, argc C.int) {
	var offset = unsafe.Sizeof(uintptr(0))
	var out []string
	for index := C.int(0); index < argc; index++ {
		out = append(out, C.GoString(*argv))
		argv = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset))
	}
	temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc,  out)

	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export RVExtension
func RVExtension(output *C.char, outputsize C.size_t, input *C.char) {
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