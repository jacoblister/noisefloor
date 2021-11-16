package cppprocessor

/*
#cgo CFLAGS: -O2
#cgo LDFLAGS: -lm

#include "processor.h"

static inline void CCall() {
}

extern void goCallback(void);

static inline void CCallCallback() {
	goCallback();
}


*/
import "C"
import "unsafe"

func CCall() {
	C.CCall()
}

//export goCallback
func goCallback() {
}

func CCallCallback() {
	C.CCallCallback()
}

func CStart() {
	C.start()
}

func CProcess() {
	const frameLength = 1024

	var samples [frameLength]float32

	C.process(frameLength, (*C.float)(unsafe.Pointer(&samples)))
}
