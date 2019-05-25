package dspUI

import (
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/stretchr/testify/assert"
)

func TestUpdateConnector(t *testing.T) {
	// Given ... Two processors, no existing connections
	gainA := processor.Gain{}
	gainB := processor.Gain{}
	connectorList := []dsp.Connector{dsp.Connector{FromProcessor: &gainA, FromPort: 0}}
	graph := dsp.Graph{ConnectorList: connectorList}
	engine := &Engine{Engine: &dsp.Engine{Graph: graph}, state: &EngineState{editState: connectionEdit}}

	// When ... No target processor
	engine.updateConnector(&engine.Engine.Graph.ConnectorList[0], true, nil, 0)

	// Then ... No connection is made
	assert.Equal(t, []dsp.Connector{}, engine.Engine.Graph.ConnectorList)

	// Given ...
	engine.Engine.Graph.ConnectorList = []dsp.Connector{dsp.Connector{FromProcessor: &gainA, FromPort: 0}}
	engine.state.selectedConnectorIsInput = true

	// When ... Valid target processor
	engine.updateConnector(&engine.Engine.Graph.ConnectorList[0], true, &gainB, 0)

	// Then ... Connection is made
	expected := []dsp.Connector{dsp.Connector{FromProcessor: &gainA, FromPort: 0, ToProcessor: &gainB, ToPort: 0}}
	assert.Equal(t, expected, engine.Engine.Graph.ConnectorList)
}
