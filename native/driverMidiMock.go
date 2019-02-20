package main

import "github.com/jacoblister/noisefloor/common/midi"

type driverMidiMock struct {
}

func (d *driverMidiMock) start() {}
func (d *driverMidiMock) stop()  {}
func (d *driverMidiMock) readEvents() []midi.Event {
	println("Mock Midi read events...")
	return nil
}
func (d *driverMidiMock) writeEvents([]midi.Event) {
	println("Mock Midi write events...")
}
