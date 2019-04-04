#include "include/platform.hpp"
#include <iostream>
#include <thread>

#if PROCESS_MOCK
#include "ProcessMock.hpp"
#define PROCESS ProcessMock
#endif
#if PROCESS_V8
#define PROCESS ProcessV8Engine
#include "ProcessV8Engine.hpp"
#endif
#if PROCESS_CPP
#define PROCESS ProcessCPP
#include "ProcessCPP.hpp"
#endif

#if AUDIO_MOCK
#include "DriverAudioMock.hpp"
#define DRIVER_AUDIO DriverAudioMock
#endif
#if AUDIO_JACK
#include "DriverAudioJack.hpp"
#define DRIVER_AUDIO DriverAudioJack
#endif
#if AUDIO_ASIO
#include "DriverAudioASIO.hpp"
#define DRIVER_AUDIO DriverAudioASIO
#endif

#if MIDI_MOCK
#include "DriverMidiMock.hpp"
#define DRIVER_MIDI DriverMidiMock
#endif
#if MIDI_WDM
#include "DriverMidiWDM.hpp"
#define DRIVER_MIDI DriverMidiWDM
#endif

#if CLIENT_MOCK
#include "ClientMock.hpp"
#define CLIENT ClientMock
#endif
#if CLIENT_RESTSERVER
#include "ClientRESTServer.hpp"
#define CLIENT ClientRESTServer
#endif

class NoiseFloor {
  public:
    NoiseFloor(bool nothing) : process(), driverAudio(process), driverMidi(), client(process) {}

    void run(void);
  private:
    PROCESS      process;
    DRIVER_AUDIO driverAudio;
    DRIVER_MIDI  driverMidi;
    CLIENT       client;
};

void NoiseFloor::run(void) {
    result<bool> result;

    driverMidi.init();
    if (!(result = driverMidi.start())) {
        std::cout << "Midi start failed: " << result.errorMessage() << std::endl;
        return;
    }

    driverAudio.init();
    driverAudio.setMidiDriver(&driverMidi);
    if (!(result = driverAudio.start())) {
        std::cout << "Audio start failed: " << result.errorMessage() << std::endl;
        return;
    }

    if(!(result = client.start())) {
        std::cout << "Client Start Failed: " << result.errorMessage() << std::endl;
    }

    // Run until ESC pressed
    std::cout << "Press ESC to exit" << std::endl;
    while (getch() != 27) { }

    client.stop();
    driverAudio.stop();
}

int main(int argc, char* argv[]) {
    NoiseFloor noiseFloor(true);

    noiseFloor.run();

    return 0;
}
