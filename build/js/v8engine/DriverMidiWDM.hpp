#include "DriverMidi.hpp"

#include <windows.h>
#include <iostream>

#define WDM_INPUT_BUFFER_SIZE 1024
#define WDM_MIDI_EVENT_SIZE   3

class DriverMidiWDM: public DriverMidi {
  private:
   char inputBuffer[WDM_INPUT_BUFFER_SIZE];
   int inputBufferIndex;

   HMIDIIN inHandle;
   HMIDIOUT outHandle;
   std::vector<MIDIEvent> midiInEvents;
  public:
    DriverMidiWDM() {}
    virtual result<bool> init();
    virtual result<bool> start();
    virtual std::vector<MIDIEvent> readMidiEvents();
    virtual void writeMidiEvents(std::vector<MIDIEvent> midiIn);
    virtual result<bool> stop();

    inline void addInputEvent(byte byte0, byte byte1, byte byte2) {
        inputBuffer[inputBufferIndex + 0] = byte0;
        inputBuffer[inputBufferIndex + 1] = byte1;
        inputBuffer[inputBufferIndex + 2] = byte2;

        MIDIEvent event = {0, WDM_MIDI_EVENT_SIZE, &inputBuffer[inputBufferIndex]};
        midiInEvents.push_back(event);

        if (inputBufferIndex + WDM_MIDI_EVENT_SIZE > (WDM_INPUT_BUFFER_SIZE - WDM_MIDI_EVENT_SIZE)) {
            inputBufferIndex = 0;
        } else {
            inputBufferIndex += WDM_MIDI_EVENT_SIZE;
        }
    }
};
