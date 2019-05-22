package synthUI

import (
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/synth"
	"github.com/jacoblister/noisefloor/app/audiomodule/synth/processor"
	"github.com/stretchr/testify/assert"
)

func TestUpdateConnector(t *testing.T) {
	// Given ... Two processors, no existing connections
	gainA := processor.Gain{}
	gainB := processor.Gain{}
	connectorList := []synth.Connector{synth.Connector{FromProcessor: &gainA, FromPort: 0}}
	graph := synth.Graph{ConnectorList: connectorList}
	engine := &Engine{Engine: &synth.Engine{Graph: graph}, state: &EngineState{}}

	// When ... No target processor
	engine.updateConnector(&engine.Engine.Graph.ConnectorList[0], true, nil, 0)

	// Then ... No connection is made
	assert.Equal(t, []synth.Connector{}, engine.Engine.Graph.ConnectorList)

	// Given ...
	engine.Engine.Graph.ConnectorList = []synth.Connector{synth.Connector{FromProcessor: &gainA, FromPort: 0}}
	engine.state.selectedConnectorIsInput = true

	// When ... Valid target processor
	engine.updateConnector(&engine.Engine.Graph.ConnectorList[0], true, &gainB, 0)

	// Then ... Connection is made
	expected := []synth.Connector{synth.Connector{FromProcessor: &gainA, FromPort: 0, ToProcessor: &gainB, ToPort: 0}}
	assert.Equal(t, expected, engine.Engine.Graph.ConnectorList)
}
