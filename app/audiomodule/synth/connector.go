package synth

//Connector specifies a connection between two Processors
type Connector struct {
	FromProcessor Processor
	FromPort      int
	ToProcessor   Processor
	ToPort        int
}
