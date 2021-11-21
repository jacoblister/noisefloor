#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "processor.h"

#define FRAME_LENGTH 48000

struct timespec clock_start={0,0}, clock_end={0,0};
void bench_start(void) {
    clock_gettime(CLOCK_MONOTONIC, &clock_start);
}

void bench_end(char *name, int times) {
    clock_gettime(CLOCK_MONOTONIC, &clock_end);
    long elapsed = 
        (clock_end.tv_sec * 1000000000 + clock_end.tv_nsec) -
        (clock_start.tv_sec * 1000000000 + clock_start.tv_nsec);

    printf("%s: %d times, %ldns, %ld ns/op (%fms, %fms/op)\n", name, times, elapsed, elapsed/times, elapsed / 1000000.0, elapsed/times/1000000.0);
}


int main(int argc, char **argv) {
    int times = 1000;
    if (argc == 2) {
        times = atoi(argv[1]);
    }
    
    float samples[FRAME_LENGTH];
 
    bench_start();
    synth_start();
    for (int i = 0; i < times; i++) {
        synth_process(FRAME_LENGTH, samples, 0, 0, 0);
    }
    bench_end("synth", times);

    bench_start();
    synthpoly_start();
    for (int i = 0; i < times; i++) {
        synthpoly_silenceprocess(FRAME_LENGTH, samples);
    }
    bench_end("polysynth", times);
}