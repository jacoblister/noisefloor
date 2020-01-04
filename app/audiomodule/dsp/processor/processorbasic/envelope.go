package processorbasic

// Phase enumerated type
type Phase int

// Basic wave shapes
const (
	Inactive Phase = iota
	Attack
	Decay
	Sustain
	Release
)

// Envelope - ADSR envelope generator
type Envelope struct {
	Attack  float32 `default:"2" min:"0" max:"100"`
	Decay   float32 `default:"100" min:"0" max:"1000"`
	Sustain float32 `default:"0.75" min:"0" max:"1"`
	Release float32 `default:"1000" min:"0" max:"1000"`

	sampleRate  float32
	output      float32
	phase       Phase
	delta       float32
	lastTrigger float32
}

// Start - init envelope generator
func (e *Envelope) Start(sampleRate int) {
	e.sampleRate = float32(sampleRate)
}

// Process - produce next sample
func (e *Envelope) Process(Gte float32, Trg float32) (Out float32) {
	if Trg > 0 && e.lastTrigger == 0 {
		e.output = 0
		e.delta = (1000 / e.Attack) / e.sampleRate
		e.phase = Attack
	}

	switch phase := e.phase; phase {
	case Attack:
		e.output += e.delta
		if e.output > 1 {
			e.delta = (1000 / e.Decay) / e.sampleRate
			e.phase = Decay
		}
	case Decay:
		e.output -= e.delta
		if e.output < e.Sustain {
			e.phase = Sustain
		}
	case Sustain:
		if Gte == 0 {
			e.delta = (1000 / e.Release) / e.sampleRate
			e.phase = Release
		}
	case Release:
		e.output -= e.delta
		if e.output < 0 {
			e.output = 0
			e.phase = Inactive
		}
	}

	e.lastTrigger = Trg
	Out = e.output
	return
}
