#pragma once

#include "../Processor/Oscillator.hpp"

class Patch {
private:
    Oscillator oscillator;

public:
    Patch(void) : oscillator() {}

    inline void start(int sampleRate) {
        oscillator.start(sampleRate);
    }

    inline float process(float freq, float gate, float trigger) {
        oscillator.freq = freq;
        float sample = oscillator.process();

        return sample;
    }
};