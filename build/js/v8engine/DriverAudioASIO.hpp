#include "Process.hpp"
#include "DriverMidi.hpp"

#define WIN32 1
#include "asio/asiosys.h"		// platform definition
#include "asio/asio.h"
#include "asio/asiodrivers.h"


class DriverAudioASIO {
  public:
    DriverAudioASIO(Process& process) : process(process) { }
    result<bool> init();
    result<bool> start();
    result<bool> stop();

    inline DriverMidi *getMidiDriver(void )            { return this->driverMidi;       }
    inline void setMidiDriver(DriverMidi *driverMidi)  { this->driverMidi = driverMidi; }

    inline Process& getProcess(void)                { return process;            }
    inline long getInputChannels(void)              { return inputChannels;      }
    inline long getOutputChannels(void)             { return outputChannels;     }
    inline long getPreferredSize(void)              { return preferredSize;      }
    inline ASIOBufferInfo*  getBufferInfo(void)     { return bufferInfo.data();  }
    inline ASIOChannelInfo* getChannelInfo(void)    { return channelInfo.data(); }
    inline bool getPostOutput(void)                 { return postOutput;         }
    inline std::vector<float *> getSamplesIn(void)  { return samplesIn;          }
    inline std::vector<float *> getSamplesOut(void) { return samplesOut;         }
  private:
    Process& process;
    DriverMidi *driverMidi = NULL;

    // Driver info
    AsioDrivers asioDrivers;
    ASIODriverInfo driverInfo;
    ASIOCallbacks asioCallbacks;
    bool postOutput;                    // PostOutput optimisation flag
    long inputChannels;
    long outputChannels;
	long minSize;
	long maxSize;
	long preferredSize;
	long granularity;
	ASIOSampleRate sampleRate;

	// Buffers
	std::vector<ASIOBufferInfo>  bufferInfo;
	std::vector<ASIOChannelInfo> channelInfo;

	std::vector<float *> samplesIn;
	std::vector<float *> samplesOut;
};
