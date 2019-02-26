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
	jack_port_t* jack_midi_input_port;
	jack_port_t* jack_midi_output_port;

	int channel_in_count;
	int channel_out_count;
	float* channel_in[MAX_CHANNELS];
	float* channel_out[MAX_CHANNELS];
} jack_c_client;

jack_c_client client;

extern int goProcess(void *arg, int blockSize,
	int channelInCount, void *channelIn,
	int channelOutCount, void *channelOut,
	int midiInCount, void *midiIn,
	int midiOutCount, void *midiOut);

#include <stdio.h>
int process_jack(jack_nframes_t nframes, void *arg) {
	printf("jack process\n");

	// static float samples1[] = {40,41,42};
	// static float samples2[] = {80,81,82};
	// client.channel_in[0] = samples1;
	// client.channel_in[1] = samples2;
	// printf("samples1 %p\n", samples1);
	// printf("samples2 %p\n", samples2);

	client.channel_in[0] = jack_port_get_buffer(client.jack_audio_input_port[0],  nframes);
	client.channel_in[1] = jack_port_get_buffer(client.jack_audio_input_port[1],  nframes);
	client.channel_out[0] = jack_port_get_buffer(client.jack_audio_output_port[0],  nframes);
	client.channel_out[1] = jack_port_get_buffer(client.jack_audio_output_port[1],  nframes);
	printf("channelin %p\n", client.channel_in);

	goProcess(arg, nframes,
		client.channel_in_count, client.channel_in,
		client.channel_out_count, client.channel_out,
		0, NULL,
		0, NULL
	);

	return 0;
}

jack_c_client* gojack_client_open(uintptr_t arg) {
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
    client.jack_midi_input_port      = jack_port_register(client.jack_client, "midi-input",  JACK_DEFAULT_MIDI_TYPE, JackPortIsInput, 0);
    client.jack_midi_output_port     = jack_port_register(client.jack_client, "midi-output", JACK_DEFAULT_MIDI_TYPE, JackPortIsOutput, 0);
	client.channel_in_count 		 = 2;
	client.channel_out_count 		 = 2;

    jack_set_process_callback(client.jack_client, process_jack, (void *)arg);
    jack_activate(client.jack_client);

    return &client;
}
*/
import "C"

import (
	"unsafe"

	"github.com/jacoblister/noisefloor/component"
)

type driverAudioJack struct {
	audioProcessor component.AudioProcessor
}

func (d *driverAudioJack) setMidiDriver(driverMidi driverMidi) {
}
func (d *driverAudioJack) setAudioProcessor(audioProcessor component.AudioProcessor) {
	d.audioProcessor = audioProcessor

	uintPtr := uintptr(unsafe.Pointer(d))
	C.gojack_client_open((_Ctype_ulong)(uintPtr))
	// println(client.in_channels)
}
func (d *driverAudioJack) start() {
}

func (d *driverAudioJack) stop() {
}
