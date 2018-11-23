package engine

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	// Process(vars ...[]*AudioFloat)
}
