#include "Process.hpp"
#include "Processor/Oscillator.hpp"
#include "Processor/MIDIInput.hpp"

class ProcessCPP: public Process {
  private:
    int samplesPerFrame;

    MIDIInput midiInput;
    Oscillator oscillator;
  public:
    ProcessCPP() : midiInput(), oscillator() {}
    virtual result<bool> init(void);
    virtual result<bool> start(int samplingRate, int samplesPerFrame);
    virtual result<bool> process(std::vector<float *> samplesIn, std::vector<float *> samplesOut, std::vector<MIDIEvent> midiIn, std::vector<MIDIEvent> midiOut);
    virtual result<bool> stop(void);
    virtual std::string query(std::string endpoint, std::string request);
};
