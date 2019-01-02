#pragma once
#include "include/result.hpp"
#include "include/midiEvent.hpp"

#include <string>
#include <vector>

class Process {
  public:
    Process() {}
    virtual result<bool> init() { return false; };
    virtual result<bool> start(int samplingRate, int samplesPerFrame) { return false; };
    virtual result<bool> process(std::vector<float *> samplesIn, std::vector<float *> samplesOut, std::vector<MIDIEvent> midiIn, std::vector<MIDIEvent> midiOut) { return false; };
    virtual result<bool> stop() { return false; };
    virtual std::string query(std::string endpoint, std::string request) { return ""; };
};