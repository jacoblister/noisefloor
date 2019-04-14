// +build windows

package nf

import (
	"C"
	"fmt"
	"unsafe"

	"github.com/jacoblister/noisefloor/component"
)
import (
	"reflect"

	"github.com/jacoblister/noisefloor/midi"
)

type driverAudioASIO struct {
	asioDriver     *ASIODriver
	audioProcessor component.AudioProcessor
	driverMidi     driverMidi
}

//export goAudioASIOCallback
func goAudioASIOCallback(arg unsafe.Pointer, blockLength C.int,
	channelInCount C.int, channelIn unsafe.Pointer,
	channelOutCount C.int, channelOut unsafe.Pointer) {
	samplesInSlice := make([][]float32, channelInCount, channelInCount)
	samplesOutSlice := make([][]float32, channelOutCount, channelOutCount)
	blockLengthInt := int(blockLength)

	for i := 0; i < int(channelInCount); i++ {
		samplesIn := indexPointer(channelIn, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesIn), Len: blockLengthInt, Cap: blockLengthInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesInSlice[i] = s
	}

	for i := 0; i < int(channelOutCount); i++ {
		samplesOut := indexPointer(channelOut, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesOut), Len: blockLengthInt, Cap: blockLengthInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesOutSlice[i] = s
	}

	dp := *(*driverAudioASIO)(arg)
	midiInSlice := dp.driverMidi.readEvents()
	midiOutSlice := make([]midi.Event, 0, 0)

	dp.audioProcessor.Process(samplesInSlice, samplesOutSlice, midiInSlice, &midiOutSlice)
	// dp.asioDriver.ASIO.OutputReady()
}

func (d *driverAudioASIO) setMidiDriver(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioASIO) setAudioProcessor(audioProcessor component.AudioProcessor) {
	d.audioProcessor = audioProcessor
}
func (d *driverAudioASIO) start() {
	println("ASIO start")

	CoInitialize(0)

	drivers, err := ListDrivers()
	if err != nil {
		panic("ASIO cannot load drivers")
	}

	d.asioDriver = drivers["GP-10"]

	err = d.asioDriver.Open()
	if err != nil {
		panic("ASIO cannot open driver")
	}

	drv := d.asioDriver.ASIO

	// getChannels
	in, out, err := drv.GetChannels()
	if err != nil {
		panic("ASIO cannot get channels")
	}
	fmt.Printf("getChannels():        %d, %d\n", in, out)

	// getBufferSize
	minSize, maxSize, preferredSize, granularity, err := drv.GetBufferSize()
	if err != nil {
		panic("ASIO cannot get buffer size")
	}
	fmt.Printf("getBufferSize():      %d, %d, %d, %d\n", minSize, maxSize, preferredSize, granularity)

	bufferDescriptors := make([]BufferInfo, 0, in+out)
	for i := 0; i < in; i++ {
		bufferDescriptors = append(bufferDescriptors, BufferInfo{
			Channel: i,
			IsInput: true,
		})
		cinfo, err := drv.GetChannelInfo(i, true)
		if err != nil {
			panic("ASIO cannot get input channel info")
		}
		fmt.Printf(" IN%-2d: active=%v, group=%d, type=%d, name=%s\n", i+1, cinfo.IsActive, cinfo.ChannelGroup, cinfo.SampleType, cinfo.Name)
	}
	for i := 0; i < out; i++ {
		bufferDescriptors = append(bufferDescriptors, BufferInfo{
			Channel: i,
			IsInput: false,
		})
		cinfo, err := drv.GetChannelInfo(i, false)
		if err != nil {
			panic("ASIO cannot get output channel info")
		}
		fmt.Printf("OUT%-2d: active=%v, group=%d, type=%d, name=%s\n", i+1, cinfo.IsActive, cinfo.ChannelGroup, cinfo.SampleType, cinfo.Name)
	}

	err = drv.CreateBuffers(bufferDescriptors, preferredSize)
	if err != nil {
		panic("ASIO cannot create buffers")
	}

	drv.SetBufferChannels(unsafe.Pointer(d), in, out, preferredSize)
	for i := 0; i < in; i++ {
		drv.SetBufferPtr(1, i, unsafe.Pointer(bufferDescriptors[i].Buffers[0]), unsafe.Pointer(bufferDescriptors[i].Buffers[1]))
	}
	for i := 0; i < out; i++ {
		drv.SetBufferPtr(0, i, unsafe.Pointer(bufferDescriptors[i+in].Buffers[0]), unsafe.Pointer(bufferDescriptors[i+in].Buffers[1]))
	}

	fmt.Printf("Output Ready %t\n", drv.OutputReady())

	err = drv.Start()
	if err != nil {
		panic("ASIO cannot start driver")
	}

	fmt.Printf("ASIO Started\n")
}

func (d *driverAudioASIO) stop() {
	println("ASIO stop")

	drv := d.asioDriver.ASIO
	err := drv.Stop()
	if err != nil {
		panic("ASIO cannot stop driver")
	}

	drv.DisposeBuffers()

	d.asioDriver.Close()

	CoUninitialize()
}

func (d *driverAudioASIO) samplingRate() int {
	return 44100
}
