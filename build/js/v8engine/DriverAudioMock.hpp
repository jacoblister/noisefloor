#include "Process.hpp"
#include "DriverMidi.hpp"

#include <thread>
#include <atomic>

class DriverAudioMock {
  public:
    DriverAudioMock(Process& process) : process(process) { }
    result<bool> init();
    result<bool> start();
    result<bool> stop();

    inline DriverMidi *getMidiDriver(void )            { return this->driverMidi;       }
    inline void setMidiDriver(DriverMidi *driverMidi)  { this->driverMidi = driverMidi; }

    inline Process& getProcess(void) { return process; }
    inline bool getStopRequest(void) { return stopRequest; }
  private:
    DriverMidi *driverMidi = NULL;
    std::thread thread;
    std::atomic<bool> stopRequest;
    Process& process;
};
