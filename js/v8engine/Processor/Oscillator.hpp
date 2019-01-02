#pragma once

#define _USE_MATH_DEFINES
#include <math.h>

class Oscillator {
private:
    float waveTable[48000];
    float currentSample;
    float sampleRate;
public:
    float freq = 220;

    inline void start(int sampleRate) {
        this->currentSample = 0;
        this->sampleRate = sampleRate;
        for (int i = 0; i < sampleRate; i++) {
            this->waveTable[i] = sin((2 * M_PI * i) / sampleRate);
        }
    }

    inline float process(void) {
        float result = waveTable[(int)currentSample];

        currentSample += freq;
        if (currentSample >= sampleRate) { currentSample -= sampleRate; }

        return result;
    }
};