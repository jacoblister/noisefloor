package processor

//Processor interface
type Processor interface {
	start(samplerate int)
	stop()
	process(args ...float32) float32
}
