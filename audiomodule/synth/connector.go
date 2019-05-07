package synth

//Connector specifies a connection between two Processors
type connector struct {
	fromProcessor Processor
	toProcessor   Processor
	fromPort      int
	toPort        int
}
