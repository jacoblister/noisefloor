#include "math.h"

struct osc {
    float waveTable[48000];
    float currentSample;
    float sampleRate;
};

static inline void osc_start(struct osc *osc, int sampleRate) {
    osc->currentSample = 0;
    osc->sampleRate = sampleRate;
    for (int i = 0; i < sampleRate; i++) {
        osc->waveTable[i] = sin((2 * M_PI * i) / sampleRate);
    }
}

static inline float osc_process(struct osc *osc, float freq) {
    float result = osc->waveTable[(int)osc->currentSample];

    osc->currentSample += freq;
    if (osc->currentSample >= osc->sampleRate) { osc->currentSample -= osc->sampleRate; }

    return result;
}

struct patch {
    struct osc osc;
};

static struct patch patch;

static inline void start() {
    osc_start(&patch.osc, 48000);
}

static inline void process(int length, float *samples) {
    for (int i = 0; i < length; i++) {
        samples[i] = osc_process(&patch.osc, 440.0);
    }
}

// #define FRAME_SIZE 1024
// int main(void) {
//     float samples[FRAME_SIZE];

//     process(FRAME_SIZE, samples);
// }