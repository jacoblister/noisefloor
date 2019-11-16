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
	graph := dsp.Graph{Connectors: connectorList}
	engine := &Engine{Engine: &dsp.Engine{Graph: graph}, state: &EngineState{editState: connectionEdit}}

	// When ... No target processor
	engine.updateConnector(&engine.Engine.Graph.Connectors[0], true, nil, 0)

	// Then ... No connection is made
	assert.Equal(t, []dsp.Connector{}, engine.Engine.Graph.Connectors)

	// Given ...
	engine.Engine.Graph.Connectors = []dsp.Connector{dsp.Connector{FromProcessor: &gainA, FromPort: 0}}
	engine.state.selectedConnectorIsInput = true

	// When ... Valid target processor
	engine.updateConnector(&engine.Engine.Graph.Connectors[0], true, &gainB, 0)

	// Then ... Connection is made
	expected := []dsp.Connector{dsp.Connector{FromProcessor: &gainA, FromPort: 0, ToProcessor: &gainB, ToPort: 0}}
	assert.Equal(t, expected, engine.Engine.Graph.Connectors)
}

func TestGetUniqueProcessorName(t *testing.T) {
	// Given ... no existing names
	existingNames := []string{}

	// When ...
	result := getUniqueProcessorName("gain", existingNames)

	// Then ... base name
	assert.Equal(t, "gain", result)

	// Given ... base name used
	existingNames = []string{"gain"}

	// When ...
	result = getUniqueProcessorName("gain", existingNames)

	// Then ... allocated 1 index
	assert.Equal(t, "gain1", result)

	// Given ... base name and index name used
	existingNames = []string{"gain", "gain2"}

	// When ...
	result = getUniqueProcessorName("gain", existingNames)

	// Then ... allocated 3 index
	assert.Equal(t, "gain3", result)
}
