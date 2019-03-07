// +build linux

package main

/*
#cgo linux LDFLAGS: -ljack
#cgo windows,386 LDFLAGS: -llibjack
#cgo windows,amd64 LDFLAGS: -llibjack64

#include <jack/jack.h>
#include <jack/midiport.h>

#define MAX_CHANNELS 8

typedef struct {
	jack_client_t* jack_client;
	jack_port_t* jack_audio_input_port[2];
	jack_port_t* jack_audio_output_port[2];

	int channel_in_count;
	int channel_out_count;
	float* channel_in[MAX_CHANNELS];
	float* channel_out[MAX_CHANNELS];
} jack_c_client;

jack_c_client client;

extern void goAudioJackCallback(void *arg, int blockSize,
	int channelInCount, void *channelIn,
	int channelOutCount, void *channelOut);

#include <stdio.h>
static inline int process_jack(jack_nframes_t nframes, void *arg) {
	client.channel_in[0] = jack_port_get_buffer(client.jack_audio_input_port[0],  nframes);
	client.channel_in[1] = jack_port_get_buffer(client.jack_audio_input_port[1],  nframes);
	client.channel_out[0] = jack_port_get_buffer(client.jack_audio_output_port[0],  nframes);
	client.channel_out[1] = jack_port_get_buffer(client.jack_audio_output_port[1],  nframes);

	goAudioJackCallback(arg, nframes,
		client.channel_in_count, client.channel_in,
		client.channel_out_count, client.channel_out
	);

	return 0;
}

static inline int gojack_sample_rate() {
	return jack_get_sample_rate(client.jack_client);
}

static inline jack_c_client* gojack_client_open(uintptr_t arg) {
    const char **ports;
    const char *client_name = "noisefloor";
    const char *server_name = NULL;
    jack_options_t options = JackNullOption;
    jack_status_t status;

    // open a client connection to the JACK server
    client.jack_client               = jack_client_open(client_name, options, &status, server_name);
    client.jack_audio_input_port[0]  = jack_port_register(client.jack_client, "input_0",  JACK_DEFAULT_AUDIO_TYPE, JackPortIsInput, 0);
    client.jack_audio_input_port[1]  = jack_port_register(client.jack_client, "input_1",  JACK_DEFAULT_AUDIO_TYPE, JackPortIsInput, 0);
    client.jack_audio_output_port[0] = jack_port_register(client.jack_client, "output_0", JACK_DEFAULT_AUDIO_TYPE, JackPortIsOutput, 0);
    client.jack_audio_output_port[1] = jack_port_register(client.jack_client, "output_1", JACK_DEFAULT_AUDIO_TYPE, JackPortIsOutput, 0);
	client.channel_in_count 		 = 2;
	client.channel_out_count 		 = 2;

    jack_set_process_callback(client.jack_client, process_jack, (void *)arg);
    jack_activate(client.jack_client);

    return &client;
}

static inline int gojack_client_sampling_rate() {
	return jack_get_sample_rate(client.jack_client);
}

*/
import "C"

import (
	"reflect"
	"unsafe"

	"github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/component"
)

type driverAudioJack struct {
	audioProcessor component.AudioProcessor
	driverMidi     driverMidi
}

//export goAudioJackCallback
func goAudioJackCallback(arg unsafe.Pointer, blockSize C.int,
	channelInCount C.int, channelIn unsafe.Pointer,
	channelOutCount C.int, channelOut unsafe.Pointer) {
	samplesInSlice := make([][]float32, channelInCount, channelInCount)
	samplesOutSlice := make([][]float32, channelOutCount, channelOutCount)
	blockSizeInt := int(blockSize)

	for i := 0; i < int(channelInCount); i++ {
		samplesIn := indexPointer(channelIn, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesIn), Len: blockSizeInt, Cap: blockSizeInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesInSlice[i] = s
	}

	for i := 0; i < int(channelOutCount); i++ {
		samplesOut := indexPointer(channelOut, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesOut), Len: blockSizeInt, Cap: blockSizeInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesOutSlice[i] = s
	}

	dp := *(*driverAudioJack)(arg)
	midiInSlice := dp.driverMidi.readEvents()
	midiOutSlice := make([]midi.Event, 0, 0)

	dp.audioProcessor.Process(samplesInSlice, samplesOutSlice, midiInSlice, &midiOutSlice)
}

func (d *driverAudioJack) setMidiDriver(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioJack) setAudioProcessor(audioProcessor component.AudioProcessor) {
	d.audioProcessor = audioProcessor
}
func (d *driverAudioJack) start() {
	uintPtr := uintptr(unsafe.Pointer(d))
	C.gojack_client_open((C.ulong)(uintPtr))
}

func (d *driverAudioJack) stop() {
}

func (d *driverAudioJack) samplingRate() int {
	println(C.gojack_client_sampling_rate())
	return int(C.gojack_client_sampling_rate())
}
