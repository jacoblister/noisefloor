// +build linux

package nf

/*
#cgo linux LDFLAGS: -ljack
#cgo windows,386 LDFLAGS: -llibjack
#cgo windows,amd64 LDFLAGS: -llibjack64

#include <stdint.h>
#include <string.h>
#include <jack/jack.h>
#include <jack/midiport.h>
typedef struct {
	jack_client_t* jack_client;
	jack_port_t* jack_midi_input_port;
	jack_port_t* jack_midi_output_port;
	uint8_t byte_buffer[1024];
} jack_midi_client;

jack_midi_client midi_client;

static inline jack_midi_client* gojack_midi_client_open(void) {
    const char **ports;
    const char *client_name = "noisefloor-midi";
    const char *server_name = NULL;
    jack_options_t options = JackNullOption;
    jack_status_t status;

    // open a client connection to the JACK server
    midi_client.jack_client               = jack_client_open(client_name, options, &status, server_name);
    midi_client.jack_midi_input_port      = jack_port_register(midi_client.jack_client, "midi-input",  JACK_DEFAULT_MIDI_TYPE, JackPortIsInput, 0);
    midi_client.jack_midi_output_port     = jack_port_register(midi_client.jack_client, "midi-output", JACK_DEFAULT_MIDI_TYPE, JackPortIsOutput, 0);

    jack_activate(midi_client.jack_client);

    return &midi_client;
}

static inline int gojack_midi_client_read(uint8_t **buffer) {
	void *midi_in_port_buf = jack_port_get_buffer(midi_client.jack_midi_input_port, 0);
	int midi_event_count = jack_midi_get_event_count(midi_in_port_buf);
	int byte_index = 0;

	for (int i = 0; i < midi_event_count; i++) {
		jack_midi_event_t read_event;
		jack_midi_event_get(&read_event, midi_in_port_buf, i);
		int32_t time = read_event.time;

		midi_client.byte_buffer[byte_index] = read_event.size;
		memcpy(&midi_client.byte_buffer[byte_index + 1], &time, 4);
		memcpy(&midi_client.byte_buffer[byte_index + 5], read_event.buffer, read_event.size);

		byte_index += read_event.size + 5;
	}

	midi_client.byte_buffer[byte_index] = 0;	// termination
	byte_index++;

	// midi_client.byte_buffer[0] = 3;
	// midi_client.byte_buffer[5] = 0x90;
	// midi_client.byte_buffer[6] = 60;
	// midi_client.byte_buffer[7] = 100;
	// midi_client.byte_buffer[8] = 0;
	// byte_index = 9;

	*buffer = midi_client.byte_buffer;
	return byte_index;
}

*/
import "C"
import (
	"unsafe"

	"github.com/jacoblister/noisefloor/midi"
)

type driverMidiJack struct {
}

func (d *driverMidiJack) start() {
	C.gojack_midi_client_open()
}

func (d *driverMidiJack) stop() {
}

func (d *driverMidiJack) readEvents() []midi.Event {
	var byteBuffer unsafe.Pointer

	byteBufferLength := C.gojack_midi_client_read((**C.uint8_t)(unsafe.Pointer(&byteBuffer)))
	cBuf := (*[1 << 30]byte)(byteBuffer)

	midiEvents := midi.DecodeByteBuffer(cBuf[:byteBufferLength])

	return midiEvents
}

func (d *driverMidiJack) writeEvents([]midi.Event) {
}
