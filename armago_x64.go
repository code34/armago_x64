package main

import "C"

import (
	"fmt"
	"os"
)

//export RVExtension
func RVExtension(output *C.char, outputsize C.int, input *C.char) {
	f, _ := os.Create("mylogs.txt")
	defer f.Close()
	var temp string
	temp = fmt.Sprintf("%s \n", C.GoString(input))
	f.WriteString(temp)
	f.Sync()
}

func main() {}