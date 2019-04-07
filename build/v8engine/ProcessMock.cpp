#include "ProcessMock.hpp"

#include <iostream>

result<bool> ProcessMock::init(void) {
    std::cout << "init process mock" << std::endl;
    return true;
}

result<bool> ProcessMock::start(int sampling_rate, int samples_per_frame) {
    return true;
}

result<bool> ProcessMock::process(std::vector<float *> samplesIn, std::vector<float *> samplesOut, std::vector<MIDIEvent> midiIn, std::vector<MIDIEvent> midiOut) {
    return true;
}

result<bool> ProcessMock::stop(void) {
    return true;
}

std::string ProcessMock::query(std::string endpoint, std::string request) {
    return "";
}
