#include <stdio.h>
#include <unistd.h>
#include <jack/jack.h>
#include <jack/midiport.h>

#include "processor.h"

jack_port_t *audio_output_port;
jack_port_t *midi_input_port;
jack_client_t *client;

float synth_freq = 0;
float synth_trigger = 0;
float synth_gate = 0;

int jack_process(jack_nframes_t nframes, void *args) {
    void *midi_in = jack_port_get_buffer(midi_input_port, 0);
    int midi_event_count = jack_midi_get_event_count(midi_in);
    
    synth_trigger = 0;
    for(int i = 0; i < midi_event_count; i++) {
        jack_midi_event_t read_event;
        jack_midi_event_get(&read_event, midi_in, i);
        unsigned char *ebuffer = ((unsigned char *)read_event.buffer);
    
        int midi_type = ebuffer[0] >> 4;
        int midi_note = ebuffer[1];
        if (midi_type == 9) {
            synth_freq = 220 * pow(2.0, (midi_note - 57) / 12.0);
            synth_trigger = 1;
            synth_gate = 1;
            printf("note on freq=%f\n", synth_freq);
        } 

        if (midi_type == 8) {
            synth_gate = 0;
            printf("note off\n");
        }
    }
    
    jack_default_audio_sample_t *out = jack_port_get_buffer(audio_output_port, nframes);

    synth_process(nframes, out, synth_freq, synth_trigger, synth_gate);
    return 0;
}

int main(void) {
    synth_start();

    client = jack_client_open("synthbench", JackNullOption, NULL);
    jack_set_process_callback(client, jack_process, 0);

    audio_output_port = jack_port_register(client, "output", JACK_DEFAULT_AUDIO_TYPE, JackPortIsOutput, 0);
    midi_input_port = jack_port_register(client, "midi-input", JACK_DEFAULT_MIDI_TYPE, JackPortIsInput, 0);

    jack_activate(client);

    sleep(-1);
}