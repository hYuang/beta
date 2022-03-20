package main

import (
	"C"
	"fmt"
)

// #cgo CFLAGS: -I{SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #inlcude <callC.h>
func main() {
	fmt.Println("A Go statement")
	C.cHello()
	fmt.Println("Another Go statement")

}
