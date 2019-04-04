#pragma once
#include "Process.hpp"

class DriverMidi {
  public:
    DriverMidi() {}
    virtual result<bool> init() { return false; };
    virtual result<bool> start() { return false; };
    virtual std::vector<MIDIEvent> readMidiEvents() { return {}; };
    virtual void writeMidiEvents(std::vector<MIDIEvent> midiIn) { };
    virtual result<bool> stop() { return false; };
};