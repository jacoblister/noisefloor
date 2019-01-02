ifeq ($(OS),Windows_NT)
    # Windows
    TARGET_PLATFORM=WINDOWS
    TARGET_AUDIO=AUDIO_ASIO
    TARGET_MIDI=MIDI_WDM
    TARGET_PROCESS=PROCESS_CPP
    TARGET_CLIENT=CLIENT_MOCK

    CPP=cl
    CFLAGS=/EHsc
    O=obj

    LINKEXTRA=User32.lib Advapi32.lib Ole32.lib Winmm.lib
else
    # Linux
    TARGET_PLATFORM=LINUX
    TARGET_AUDIO=AUDIO_JACK
    TARGET_MIDI=MIDI_MOCK
    TARGET_PROCESS=PROCESS_CPP
    TARGET_CLIENT=CLIENT_MOCK

    CPP=g++
    V8_DIR=/home/jacob/v8
    CFLAGS=-I$(V8_DIR) -I$(V8_DIR)/include -std=c++0x
    O=o

    LINKFLAGS=-Wl,--start-group                                                     \
        $(V8_DIR)/out.gn/x64.release/obj/libv8_base.a                               \
        $(V8_DIR)/out.gn/x64.release/obj/libv8_libbase.a                            \
        $(V8_DIR)/out.gn/x64.release/obj/libv8_external_snapshot.a                  \
        $(V8_DIR)/out.gn/x64.release/obj/libv8_libplatform.a                        \
        $(V8_DIR)/out.gn/x64.release/obj/libv8_libsampler.a                         \
        $(V8_DIR)/out.gn/x64.release/obj/third_party/icu/libicuuc.a                 \
        $(V8_DIR)/out.gn/x64.release/obj/third_party/icu/libicui18n.a               \
        -Wl,--end-group -lrt -ldl -pthread -std=c++0x -lc++ -ljack

    LINKEXTRA=-o NoiseFloor
endif
LINK=$(CPP)
DEFINES=-D$(TARGET_PLATFORM) -D$(TARGET_AUDIO) -D$(TARGET_MIDI) -D$(TARGET_PROCESS) -D$(TARGET_CLIENT)

%.o: %.cpp
	$(CPP) $(DEFINES) -c -O2 -o $@ $< $(CFLAGS)

%.obj: %.cpp
	$(CPP) $(DEFINES) -c -O2 $< $(CFLAGS)

all: NoiseFloor

ifeq ($(TARGET_AUDIO), AUDIO_MOCK)
    PLATFORM_DRIVER_AUDIO = DriverAudioMock.$O
endif
ifeq ($(TARGET_AUDIO), AUDIO_JACK)
    PLATFORM_DRIVER_AUDIO = DriverAudioJack.$O
endif
ifeq ($(TARGET_AUDIO), AUDIO_ASIO)
    PLATFORM_OBJ = asio.obj asiolist.obj asiodrivers.obj
    PLATFORM_DRIVER_AUDIO = DriverAudioASIO.$O
endif

ifeq ($(TARGET_MIDI), MIDI_MOCK)
    PLATFORM_DRIVER_MIDI = DriverMidiMock.$O
endif
ifeq ($(TARGET_MIDI), MIDI_WDM)
    PLATFORM_DRIVER_MIDI = DriverMidiWDM.$O
endif

ifeq ($(TARGET_PROCESS), PROCESS_MOCK)
    PLATFORM_PROCESS = ProcessMock.$O
endif
ifeq ($(TARGET_PROCESS), PROCESS_V8)
    PLATFORM_PROCESS = ProcessV8Engine.$O
endif
ifeq ($(TARGET_PROCESS), PROCESS_CPP)
    PLATFORM_PROCESS = ProcessCPP.$O
endif

ifeq ($(TARGET_CLIENT), CLIENT_MOCK)
    PLATFORM_CLIENT = ClientMock.$O
endif
ifeq ($(TARGET_CLIENT), CLIENT_RESTSERVER)
    PLATFORM_CLIENT = ClientRESTServer.$O
endif


asio.obj: asio/asio.cpp
	cl -c -DWIN32 asio/asio.cpp

asiodrivers.obj: asio/asiodrivers.cpp
	cl -c -DWIN32 asio/asiodrivers.cpp

asiolist.obj: asio/asiolist.cpp
	cl -c -DWIN32 asio/asiolist.cpp

NoiseFloor.$O: NoiseFloor.cpp

ProcessMock.$O: ProcessMock.cpp ProcessMock.hpp
ProcessV8Engine.$O: ProcessV8Engine.cpp ProcessV8Engine.hpp
ProcessCPP.$O: ProcessCPP.cpp ProcessCPP.hpp Processor/MIDIInput.hpp Processor/Oscillator.hpp Test/Patch.hpp Test/MultiPatch.hpp

DriverAudioMock.$O: DriverAudioMock.cpp DriverAudioMock.hpp
DriverAudioASIO.$O: DriverAudioASIO.cpp DriverAudioASIO.hpp
DriverAudioJack.$O: DriverAudioJack.cpp DriverAudioJack.hpp

DriverMidiMock.$O: DriverMidiMock.cpp DriverMidiMock.hpp
DriverMidiWDM.$O: DriverMidiWDM.cpp DriverMidiWDM.hpp

ClientMock.$O: ClientMock.cpp ClientMock.hpp
ClientRESTServer.$O: ClientRESTServer.cpp ClientRESTServer.hpp

DRIVERS = $(PLATFORM_DRIVER_AUDIO) $(PLATFORM_DRIVER_MIDI)
PROCESS = $(PLATFORM_PROCESS)
CLIENTS = $(PLATFORM_CLIENT)
ALL = NoiseFloor.$O $(DRIVERS) $(PROCESS) $(CLIENTS) $(PLATFORM_OBJ)

NoiseFloor: $(ALL)
	$(LINK) $(ALL) $(LINKFLAGS) $(LINKEXTRA)

clean:
	rm -f *.o *.obj NoiseFloor