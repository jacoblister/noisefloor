// +build windows

package nf

/*
#cgo windows LDFLAGS: -lwinmm

#include <stdint.h>
#include <windows.h>
#include <stdio.h>

typedef struct {
    HMIDIIN hmidiin;
    HMIDIOUT hmidiout;
    int byte_index;
	uint8_t byte_buffer[1024];
} wdm_midi_client;

void CALLBACK midiCallback(HMIDIIN handle, UINT inputStatus, DWORD_PTR instancePtr, DWORD_PTR midiMessage, DWORD timestamp) {
    wdm_midi_client *client = (wdm_midi_client *)instancePtr;
    int event_size = 3;

    switch (inputStatus) {
    case MIM_DATA:
        client->byte_buffer[client->byte_index] = event_size;
        client->byte_buffer[client->byte_index + 1] = 0;
        client->byte_buffer[client->byte_index + 2] = 0;
        client->byte_buffer[client->byte_index + 3] = 0;
        client->byte_buffer[client->byte_index + 4] = 0;
        client->byte_buffer[client->byte_index + 5] = (byte)(midiMessage & 0xFF);
        client->byte_buffer[client->byte_index + 6] = (byte)((midiMessage >> 8)  & 0xFF);
        client->byte_buffer[client->byte_index + 7] = (byte)((midiMessage >> 16) & 0xFF);
        client->byte_index += event_size + 5;
        break;
    }
}

static inline wdm_midi_client* gowdm_midi_client_open(char *device_name) {
    // Get Input device number
    int inputDeviceId = -1;
    int numDevs = midiInGetNumDevs();
    MIDIINCAPS mi_caps;
    for (int i = 0; i < numDevs; i++) {
        if (!midiInGetDevCaps(i, &mi_caps, sizeof(MIDIINCAPS))) {
            if (strcmp(mi_caps.szPname, device_name) == 0) {
                inputDeviceId = i;
                break;
            }
        }
    }

    if (inputDeviceId < 0) {
        return NULL;
    }

    wdm_midi_client *client = calloc(1, sizeof(wdm_midi_client));
    int inputOpenResult = midiInOpen(&client->hmidiin, inputDeviceId, (DWORD_PTR)midiCallback, (DWORD_PTR)client, CALLBACK_FUNCTION);
    if (inputOpenResult != MMSYSERR_NOERROR ) {
        free(client);
        return NULL;
    }
    printf("WDM Open Ok\n");

    int inputStartResult = midiInStart(client->hmidiin);
    if (inputStartResult != MMSYSERR_NOERROR ) {
        free(client);
        return NULL;
    }
    printf("WDM Start Ok\n");

    return client;
}

static inline void gowdm_midi_client_close(wdm_midi_client* client) {
    printf("closing WDM\n");

    free(client);
}

static inline int gowdm_midi_client_read(wdm_midi_client* client, uint8_t **buffer) {
    int length = client->byte_index + 1;

    *buffer = client->byte_buffer;
    client->byte_buffer[client->byte_index] = 0;    // Terminiation byte
    client->byte_index = 0;
    return length;
}

*/
import "C"
import (
	"unsafe"

	"github.com/jacoblister/noisefloor/midi"
)

type driverMidiWDM struct {
	client *C.wdm_midi_client
}

func (d *driverMidiWDM) start() {
	deviceName := "MPKmini2"
	// deviceName := "Seaboard Block"

	d.client = C.gowdm_midi_client_open(C.CString(deviceName))
	if d.client == nil {
		panic("Could not open MIDI device")
	}
}

func (d *driverMidiWDM) stop() {
	C.gowdm_midi_client_close(d.client)
}

func (d *driverMidiWDM) readEvents() []midi.Event {
	var byteBuffer unsafe.Pointer

	byteBufferLength := C.gowdm_midi_client_read(d.client, (**C.uint8_t)(unsafe.Pointer(&byteBuffer)))
	cBuf := (*[1 << 30]byte)(byteBuffer)

	midiEvents := midi.DecodeByteBuffer(cBuf[:byteBufferLength])

	return midiEvents
}

func (d *driverMidiWDM) writeEvents([]midi.Event) {
}
