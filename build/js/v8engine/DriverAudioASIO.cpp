#include "DriverAudioASIO.hpp"

#include <stdint.h>
#include <iostream>
#include <thread>

#define ASIO_DRIVER_NAME "ASIO4ALL v2"
//#define ASIO_DRIVER_NAME "GP-10"

#define SIGNED32_MAX 2147483647.0;

// Driver object, no callback parameter exists for this
static DriverAudioASIO *driver;

void ASIO_bufferSwitch(long index, ASIOBool processNow) {
    for (int i = 0; i < driver->getInputChannels(); i++) {
        for (int j = 0; j < driver->getPreferredSize(); j++) {
            driver->getSamplesIn()[i][j] =
                ((int32_t *)driver->getBufferInfo()[i].buffers[index])[j] / SIGNED32_MAX
        }
    }

    std::vector<MIDIEvent> midiIn;
    std::vector<MIDIEvent> midiOut;
    if (driver->getMidiDriver()) {
        midiIn = driver->getMidiDriver()->readMidiEvents();
    }
    driver->getProcess().process(driver->getSamplesIn(), driver->getSamplesOut(), midiIn, midiOut);

    for (int i = 0; i < driver->getOutputChannels(); i++) {
        for (int j = 0; j < driver->getPreferredSize(); j++) {
            ((int32_t *)driver->getBufferInfo()[driver->getInputChannels() + i].buffers[index])[j] =
                driver->getSamplesOut()[i][j] * SIGNED32_MAX;
        }
    }

    if (driver->getPostOutput()) {
        ASIOOutputReady();
    }
}

void ASIO_sampleRateChanged(ASIOSampleRate sRate) {
    std::cout << "ASIO_sampleRateChanged" << std::endl;
}

long ASIO_asioMessages(long selector, long value, void* message, double* opt) {
    std::cout << "ASIO_asioMessages" << std::endl;

    return 0;
}

ASIOTime *ASIO_bufferSwitchTimeInfo(ASIOTime *timeInfo, long index, ASIOBool processNow) {
    std::cout << "ASIO_bufferSwitchTimeInfo" << std::endl;

    return 0;
}

result<bool> DriverAudioASIO::init() {
    return true;
}

result<bool> DriverAudioASIO::start() {
	ASIOError res;

    // Basic init
    driver = this;
    if (!this->asioDrivers.loadDriver(ASIO_DRIVER_NAME)) {
        return result<bool>(false, "ASIO driver load failed");
    }
    if (ASIOInit(&this->driverInfo) != ASE_OK) {
        return result<bool>(false, "ASIO init failed");
    }

    ASIOControlPanel();

    // Get driver info
    if (ASIOGetChannels(&this->inputChannels, &this->outputChannels) != ASE_OK) {
        return result<bool>(false, "ASIO get channels failed");
    }
    if (ASIOGetBufferSize(&this->minSize, &this->maxSize, &this->preferredSize, &this->granularity) != ASE_OK) {
        return result<bool>(false, "ASIO get buffer size failed");
    }
    if (ASIOGetSampleRate(&this->sampleRate) != ASE_OK) {
        return result<bool>(false, "ASIO get sample rate failed");
    }
    this->postOutput = (ASIOOutputReady() == ASE_OK);
    std::cout << "Post Output: "    << this->postOutput    << std::endl;
    std::cout << "Preferred Size: " << this->preferredSize << std::endl;

    // Allocate buffers
    this->bufferInfo.resize(this->inputChannels + this->outputChannels);
    this->samplesIn.resize(this->inputChannels);
    this->samplesOut.resize(this->outputChannels);
    ASIOBufferInfo *info = this->bufferInfo.data();
    for(int i = 0; i < this->inputChannels; i++, info++) {
        info->isInput = ASIOTrue;
        info->channelNum = i;
        info->buffers[0] = info->buffers[1] = 0;
        this->samplesIn[i] = (float *)malloc(this->preferredSize * sizeof(float));
    }
	for(int i = 0; i < this->outputChannels; i++, info++) {
		info->isInput = ASIOFalse;
		info->channelNum = i;
		info->buffers[0] = info->buffers[1] = 0;
		this->samplesOut[i] = (float *)malloc(this->preferredSize * sizeof(float));
	}
    this->asioCallbacks.bufferSwitch         = &ASIO_bufferSwitch;
    this->asioCallbacks.sampleRateDidChange  = &ASIO_sampleRateChanged;
    this->asioCallbacks.asioMessage          = &ASIO_asioMessages;
    this->asioCallbacks.bufferSwitchTimeInfo = &ASIO_bufferSwitchTimeInfo;
	res = ASIOCreateBuffers(this->bufferInfo.data(),
		this->inputChannels + this->outputChannels,
		this->preferredSize, &this->asioCallbacks);
    std::cout << "Allocating Buffers: " << res << std::endl;
    if (res != ASE_OK) {
        return result<bool>(false, "ASIO allocate buffers failed");
    }

    // Get buffer details
    this->channelInfo.resize(this->inputChannels + this->outputChannels);
    for (int i = 0; i < this->inputChannels + this->outputChannels; i++) {
        this->channelInfo[i].channel = this->bufferInfo[i].channelNum;
        this->channelInfo[i].isInput = this->bufferInfo[i].isInput;
        res = ASIOGetChannelInfo(&this->channelInfo[i]);
        if (res != ASE_OK) {
            return result<bool>(false, "ASIO get channel info failed");
        }
        if (this->channelInfo[i].type != ASIOSTInt32LSB) {
            return result<bool>(false, "ASIO 32 bit sample supported only");
        }
    }

    // Start Audio
    if (ASIOStart() != ASE_OK)  {
        return result<bool>(false, "ASIO start failed");
    }

    driver->getProcess().start(this->sampleRate, this->preferredSize);

    return true;
}

result<bool> DriverAudioASIO::stop() {
    std::cout << "ASIO Stop " << std::endl;
    ASIOStop();
    std::cout << "ASIO ASIODisposeBuffers " << std::endl;
    ASIODisposeBuffers();
//    std::cout << "ASIO ASIOExit " << std::endl;
//    ASIOExit();

    for(int i = 0; i < this->inputChannels; i++) {
        free(this->samplesIn[i]);
    }
    for(int i = 0; i < this->outputChannels; i++) {
        free(this->samplesOut[i]);
    }

    this->asioDrivers.removeCurrentDriver();

    return true;
}
