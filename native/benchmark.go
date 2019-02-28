package main

/*
static inline int c_call(void) {
    return 1;
}

extern void goCallback(void);
static inline int c_callback(void) {
    goCallback();
    return 1;
}
*/
import "C"

func makeGoCall() {
}

func makeCCall() {
	C.c_call()
}

//export goCallback
func goCallback() {
}

func makeCCallBack() {
	C.c_callback()
}
