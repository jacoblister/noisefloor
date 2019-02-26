package main

import (
	"github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/component"
)

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

type driverMidi interface {
	start()
	stop()
	readEvents() []midi.Event
	writeEvents([]midi.Event)
}

type driverAudio interface {
	setMidiDriver(driverMidi driverMidi)
	setAudioProcessor(audioProcessor component.AudioProcessor)
	start()
	stop()
}

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
