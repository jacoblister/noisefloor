#include "math.h"

/* Oscillator */

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

static inline float osc_process(struct osc *osc, float frq) {
    float result = osc->waveTable[(int)osc->currentSample];

    osc->currentSample += frq;
    if (osc->currentSample >= osc->sampleRate) { osc->currentSample -= osc->sampleRate; }

    return result;
}

/* Envelope */

struct env {
    float Attack;
    float Decay;
    float Sustain;
    float Release;

    float sampleRate;
    float output;
    int step;
    float delta;
    float lastTrigger;
};

static inline void env_start(struct env *env, int sampleRate) {
    env->sampleRate = sampleRate;
}

static inline float env_process(struct env *env, float trg, float gte) {
    if (trg > 0 && env->lastTrigger == 0) {
        env->output = 0;
        env->delta = (1000 / env->Attack) / env->sampleRate;
        env->step = 1;
    }

    env->output += env->delta;

    switch (env->step) {
    case 1:
        if (env->output > 1) {
            env->delta = (1000 / -env->Decay) / env->sampleRate;
            env->step = 2;
        }
        break;
    case 2:
        if (env->output < env->Sustain) {
            env->delta = 0;
            env->step = 3;
        }
        break;
    case 3:
        if (gte == 0) {
            env->delta = (1000 / -env->Release) / env->sampleRate;
            env->step = 4;
        }
        break;
    case 4:
        if (env->output < 0) {
            env->output = 0;
            env->delta = 0;
            env->step = 0;
        }
        break;
    }

    env->lastTrigger = 0;
    return env->output;
}

/* Gain */

struct gai {
};

static inline void gai_start(struct gai *gai, int sampleRate) {
}

static inline float gai_process(struct gai *gai, float in, float gain) {
    return in * gain;
}


/* Patch */

struct patch {
    struct osc osc;
    struct env env;
    struct gai gai;
};

static struct patch patch;

static inline void synth_start() {
    osc_start(&patch.osc, 48000);

    patch.env.Attack = 2;
    patch.env.Decay = 100;
    patch.env.Sustain = 0.75;
    patch.env.Release = 1000;
    env_start(&patch.env, 48000);

    gai_start(&patch.gai, 48000);
}

static inline void synth_process(int length, float *samples, float freq, float trigger, float gate) {
    for (int i = 0; i < length; i++) {
        float osc = osc_process(&patch.osc, freq);
        float env = env_process(&patch.env, trigger, gate);
        float out = gai_process(&patch.gai, osc, env);
        
        samples[i] = out;
    }

}

// #define FRAME_SIZE 1024
// int main(void) {
//     float samples[FRAME_SIZE];

//     process(FRAME_SIZE, samples);
// }