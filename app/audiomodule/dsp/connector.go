package dsp

//Connector specifies a connection between two Processors
type Connector struct {
	FromProcessor Processor
	FromPort      int
	ToProcessor   Processor
	ToPort        int
}

//Processor is the getter for the Connector Processor
func (c *Connector) Processor(isInput bool) Processor {
	if isInput {
		return c.ToProcessor
	}
	return c.FromProcessor
}

//Port is the getter for the Connector Port
func (c *Connector) Port(isInput bool) int {
	if isInput {
		return c.ToPort
	}
	return c.FromPort
}

//SetProcessor is the setter for the Connector Processor
func (c *Connector) SetProcessor(isInput bool, processor Processor) {
	if isInput {
		c.ToProcessor = processor
		return
	}
	c.FromProcessor = processor
}

//SetPort is the setter for the Connector Port
func (c *Connector) SetPort(isInput bool, port int) {
	if isInput {
		c.ToPort = port
		return
	}
	c.FromPort = port
}
