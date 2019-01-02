#include "DriverMidi.hpp"

class DriverMidiMock: public DriverMidi {
  public:
    DriverMidiMock() {}
    virtual result<bool> init();
    virtual result<bool> start();
    virtual std::vector<MIDIEvent> readMidiEvents();
    virtual void writeMidiEvents(std::vector<MIDIEvent> midiIn);
    virtual result<bool> stop();
};
