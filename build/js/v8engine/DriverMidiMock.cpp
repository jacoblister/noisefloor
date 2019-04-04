#include "DriverMidiMock.hpp"

result<bool> DriverMidiMock::init(void) {
    return true;
}

result<bool> DriverMidiMock::start(void) {
    return true;
}

std::vector<MIDIEvent> DriverMidiMock::readMidiEvents(void) {
    return {};
}

void DriverMidiMock::writeMidiEvents(std::vector<MIDIEvent> midiIn) {
}

result<bool> DriverMidiMock::stop(void) {
    return true;
}
