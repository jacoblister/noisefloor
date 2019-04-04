#include "Process.hpp"
#include "DriverMidi.hpp"

#include <jack/jack.h>
#include <jack/midiport.h>

class DriverAudioJack {
  public:
    DriverAudioJack(Process& process) : process(process) { }
    bool init();
    bool start();
    bool stop();

    inline void setMidiDriver(DriverMidi *driverMidi)  { }

    inline Process&       getProcess(void)         { return process; }
    inline jack_client_t* getJackClient(void)      { return jack_client; }
    inline jack_port_t*   getJackAudioInputPort(int i)  { return jack_audio_input_port[i]; }
    inline jack_port_t*   getJackAudioOutputPort(int i) { return jack_audio_output_port[i]; }
    inline jack_port_t*   getJackMIDIInputPort(void)  { return jack_midi_input_port; }
    inline jack_port_t*   getJackMIDIOutputPort(void) { return jack_midi_output_port; }
  private:
    Process& process;

    jack_client_t* jack_client;
    jack_port_t* jack_audio_input_port[2];
    jack_port_t* jack_audio_output_port[2];
    jack_port_t* jack_midi_input_port;
    jack_port_t* jack_midi_output_port;
};
