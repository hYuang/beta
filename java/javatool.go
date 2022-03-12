package main

// #cgo CFLAGS: -Iinclude -Iinclude/win32
//#include "jvmti.h"
//#include "jni.h"
//#include "jni_md.h"
import "C"
import "fmt"

//export Info
func Info() {
	fmt.Println("jvmti")
}

//export Agent_OnLoad
func Agent_OnLoad(vm *C.JavaVM, options *C.char, reserved *C.void) C.jint {
	fmt.Println("Agent load ")
	return 0
}

//export Agent_OnUnload
func Agent_OnUnload(vm *C.JavaVM) {
	fmt.Println("Agent unload")
}

func main() {

}
