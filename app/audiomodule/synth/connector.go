package synth

//Connector specifies a connection between two Processors
type Connector struct {
	FromProcessor Processor
	ToProcessor   Processor
	FromPort      int
	ToPort        int
}
