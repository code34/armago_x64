package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

//export RVExtension
func RVExtension(output *C.char, outputsize C.int, input *C.char) {
	// Create a file and write input from Arma
	f, _ := os.Create("mylogs.txt")
	defer f.Close()
	temp := fmt.Sprintf("%s", C.GoString(input))
	f.WriteString(temp)
	f.Sync()

	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func main() {}