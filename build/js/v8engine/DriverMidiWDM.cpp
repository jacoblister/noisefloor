#include "DriverMidiWDM.hpp"

#define WDM_DRIVER_NAME "MPKmini2"

void CALLBACK midiCallback(HMIDIIN handle, UINT inputStatus, DWORD_PTR instancePtr, DWORD_PTR midiMessage, DWORD timestamp) {
    DriverMidiWDM *driver = (DriverMidiWDM *)instancePtr;

    switch (inputStatus) {
    case MIM_DATA:
//        std::cout << "Got MIDI events" << std::endl;
//(byte)(dwParam1 & 0xFF), (byte)((dwParam1 >> 8) & 0xFF), (byte)((dwParam1 >> 16) & 0xFF)

//        printf("%d, %d, %d\n", (int)(midiMessage & 0xFF), (int)((midiMessage >> 8)  & 0xFF), (int)((midiMessage >> 16) & 0xFF));
//        std::cout << (byte)(midiMessage & 0xFF) << std::endl;
//        std::cout << (byte)((midiMessage >> 8)  & 0xFF) << std::endl;
//        std::cout << (byte)((midiMessage >> 16) & 0xFF) << std::endl;
        driver->addInputEvent((midiMessage & 0xFF),((midiMessage >> 8)  & 0xFF), ((midiMessage >> 16) & 0xFF));
        break;
    }
}

result<bool> DriverMidiWDM::init(void) {
    return true;
}

result<bool> DriverMidiWDM::start(void) {
    // Get Input device number
    int inputDeviceId = -1;
    int numDevs = midiInGetNumDevs();
    MIDIINCAPS mi_caps;
    for (int i = 0; i < numDevs; i++) {
        if (!midiInGetDevCaps(i, &mi_caps, sizeof(MIDIINCAPS))) {
            if (strcmp(mi_caps.szPname, WDM_DRIVER_NAME) == 0) {
                inputDeviceId = i;
                break;
            }
        }
    }
    if (inputDeviceId < 0) {
        return result<bool>(false, "WDM MIDI input device not found");
    }

    int inputOpenResult = midiInOpen(&this->inHandle, inputDeviceId, (DWORD_PTR)midiCallback, (DWORD_PTR)this, CALLBACK_FUNCTION);
    if (inputOpenResult != MMSYSERR_NOERROR ) {
        return result<bool>(false, "WDM MIDI open device failed:");
    }

    int inputStartResult = midiInStart(this->inHandle);
    if (inputStartResult != MMSYSERR_NOERROR ) {
        return result<bool>(false, "WDM MIDI start device failed:");
    }

    inputBufferIndex = 0;

    return true;
}

std::vector<MIDIEvent> DriverMidiWDM::readMidiEvents(void) {
    std::vector<MIDIEvent> events = this->midiInEvents;
    this->midiInEvents.clear();
    return events;
}

void DriverMidiWDM::writeMidiEvents(std::vector<MIDIEvent> midiIn) {
}

result<bool> DriverMidiWDM::stop(void) {
    std::cout << "WDM MIDI close" << std::endl;

    midiInStop(this->inHandle);
    midiInClose(this->inHandle);

    return true;
}
