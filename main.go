package main

import (
	"fmt"
	// #cgo LDFLAGS: -L./nativelib -lnative
	// #include <stdlib.h>
	// #include "nativelib/native.h"
	"C"
	"unsafe"
)

func main() {
	fmt.Println("Start")
	// Turn a Go string into a C string
	hello := C.CString("Hello")
	// Call a C function, convert the return value into a Go string
	valueFromCSpace := C.GoString(C.some_function(32, hello))
	// Free the C string created earlier
	C.free(unsafe.Pointer(hello))
	// Use the result from the C function.
	fmt.Println(valueFromCSpace)
	// Note that there is still an unfree'd value floating about, as some_function
	// did heap allocation
}
