package cppprocessor

/*
#cgo CFLAGS: -O2
#cgo LDFLAGS: -lm

#include "../processor.h"

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

const frameLength = 48000

func CStart() {
	C.synth_start()
}

func CProcess() {
	var samples [frameLength]float32

	C.synth_process(frameLength, (*C.float)(unsafe.Pointer(&samples)), 0, 0, 0)
}

func CPolyStart() {
	C.synthpoly_start()
}

func CPolyProcess() {
	var samples [frameLength]float32

	C.synthpoly_silenceprocess(frameLength, (*C.float)(unsafe.Pointer(&samples)))
}
