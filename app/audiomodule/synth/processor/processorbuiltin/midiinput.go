package processorbuiltin

//MIDIInput is the MIDI to CV Converter
type MIDIInput struct{}

// Start - init envelope generator
func (t *MIDIInput) Start(sampleRate int) {}

// Process - produce next sample
func (t *MIDIInput) Process() (frequency float32, gate float32, trigger float32) {
	return
}

// Definition exports processor definition
func (t *MIDIInput) Definition() (name string, inputs []string, outputs []string) {
	return "MIDIInput", []string{}, []string{"frequency", "level", "trigger"}
}

//ProcessArray calls process with an array of input/output samples
func (t *MIDIInput) ProcessArray(in []float32) (output []float32) {
	return
}
