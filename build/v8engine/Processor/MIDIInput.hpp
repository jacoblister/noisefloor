#include "../include/midiEvent.hpp"

#include <vector>
#include <array>
#include <unordered_map>

class MIDIInput {
 public:
    bool polyPhonic;
    int channels = 4;

    std::vector<int> channelNotes;
    std::vector<std::array<float, 3>> channelData;
    std::unordered_map<int, int> noteChannels;
    int nextChannel;
    int triggerClear;
 public:
    inline void start(int sampleRate) {
        this->channelNotes.resize(this->channels, 0);
        this->channelData.resize(this->channels, {0,0,0});
        this->noteChannels.clear();
        this->nextChannel = 0;
    }

    inline void processMidi(std::vector<MIDIEvent> midiInput) {
        for (int i = 0; i < midiInput.size(); i++) {
            struct MIDIEvent& midiEvent = midiInput.at(i);
            int note     = midiEvent.data[1];
            int velocity = midiEvent.data[2];

            // note release or new not - free allocated channel
            if (this->noteChannels.count(note) != 0) {
                int noteChannel = this->noteChannels.at(note);
                this->channelNotes[noteChannel] = 0;
                this->channelData[noteChannel][1] = 0;
                this->channelData[noteChannel][2] = 0;
                this->noteChannels.erase(note);
            }

            if (velocity > 0) {
                // Calculate frequency and level for note
                float frequency = 220.0 * pow(2.0, (float)(note - 57) / 12);
                float level = (float)velocity / 127;

                // Allocate next free channel
                int targetChannel = this->nextChannel;
                while (this->channelNotes[targetChannel] != 0) {
                    targetChannel++;
                    if (targetChannel >= this->channels) { targetChannel = 0; }

                    // If all channels active use current target
                    if (targetChannel == this->nextChannel) {
                        this->channelNotes[targetChannel] = 0;
                        noteChannels.erase(this->channelNotes[targetChannel]);
                    }
                }

                // set next channel, round robin
                this->nextChannel = targetChannel + 1;
                if (this->nextChannel >= this->channels) { this->nextChannel = 0; }

                // set channel active
                this->channelNotes[targetChannel] = note;
                this->channelData[targetChannel][0] = frequency;
                this->channelData[targetChannel][1] = level;
                this->channelData[targetChannel][2] = level;
                this->noteChannels[note] = targetChannel;
            }
        }

        this->triggerClear = 2;
    }

    inline std::vector<std::array<float, 3>>& process() {
        if (this->triggerClear > 0) {
            this->triggerClear--;
            if (this->triggerClear == 0) {
                // Clear triggers
                for (int i = 0; i < this->channels; i++) {
                    this->channelData[i][2] = 0;
                }
            }
        }
        return this->channelData;
    }
};