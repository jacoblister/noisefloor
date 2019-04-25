package app

import "github.com/jacoblister/noisefloor/component"

// Driver is the Audio/MIDI device drive API
type Driver interface {
	AudioDeviceList() []AudioDevice
	MIDIDeviceList() []MIDIDevice

	Start(hardwareDevices HardwareDevices, audioProcessor component.AudioProcessor)
	Stop(hardwareDevices HardwareDevices)
}

// AudioDevice is an Audio device definition
type AudioDevice struct {
	Name string
}

// MIDIDevice is a MIDI device definition
type MIDIDevice struct {
	Name string
}

// HardwareDevices is the configured Audio/MIDI devices
type HardwareDevices struct {
	audioDevice AudioDevice
	midiDevice  []MIDIDevice
}
