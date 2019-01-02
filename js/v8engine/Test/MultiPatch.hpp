#include "Patch.hpp"

#include <vector>

class MultiPatch {
private:
    std::vector<Patch> patches;

public:
    int channels = 4;

    MultiPatch(void) : patches() {}

    inline void start(int sampleRate) {
        this->patches.resize(this->channels);
    }

    inline float process(std::vector<std::array<float, 3>> freqs) {
        float result = 0;
        for (int i = 0; i < channels; i++) {
            result += patches[i].process(freqs[i][0], freqs[i][1], freqs[i][2]);
        }

        return result;
    }
};