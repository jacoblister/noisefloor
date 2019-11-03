package dsp

import (
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/stretchr/testify/assert"
)

//func TestCompileGraphExecutor_Empty(t *testing.T) {
//	// Given ... Empty graph
//	graph := Graph{}
//
//	// When ...
//	result := compileGraphExecutor(graph)
//
//	// Then ...
//	assert.Equal(t, graphExecutor{}, result)
//}
//
//func TestCompileGraphExecutor_TwoProcessors(t *testing.T) {
//	// Given ... Oscillator and output
//	graph := Graph{}
//	osc := processor.Oscillator{}
//	graph.ProcessorList = append(graph.ProcessorList,
//		ProcessorDefinition{Processor: &osc})
//	outputTerminal := processorbuiltin.Terminal{}
//	outputTerminal.SetParameters(true, 2)
//	graph.ProcessorList = append(graph.ProcessorList,
//		ProcessorDefinition{Processor: &outputTerminal})
//	graph.ConnectorList = append(graph.ConnectorList,
//		Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0})
//
//	// When ...
//	result := compileGraphExecutor(graph)
//
//	// Then ... expected 'specical' processors, and ops
//	assert.Equal(t, &outputTerminal, result.outputTerm)
//	assert.Equal(t, &osc, result.ops[0].processor)
//	assert.Equal(t, []*Connector{}, result.ops[0].connectorIn)
//	assert.Equal(t,
//		[]*Connector{
//			&Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0},
//		},
//		result.ops[0].connectorOut)
//}

func TestCompileGraphExecutor_SimplePatch(t *testing.T) {
	// Given ... Simple patch
	graph := Graph{}
	midiInput := processorbuiltin.MIDIInput{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 80, Y: 80, Processor: &midiInput})
	osc := processor.Oscillator{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 240, Y: 80, Processor: &osc})
	env := processor.Envelope{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 240, Y: 240, Processor: &env})
	gain := processor.Gain{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 400, Y: 80, Processor: &gain})
	outputTerminal := processorbuiltin.Terminal{}
	outputTerminal.SetParameters(true, 2)
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 560, Y: 80, Processor: &outputTerminal})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 0, ToProcessor: &osc, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 1, ToProcessor: &env, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 2, ToProcessor: &env, ToPort: 1})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0})

	// When ...
	result := compileGraphExecutor(graph)

	// Then ... expected 'speical' processors, and ops
	assert.Equal(t, &midiInput, result.midiInput)
	assert.Equal(t, &outputTerminal, result.outputTerm)

	assert.Equal(t, &midiInput, result.ops[0].processor)
	assert.Equal(t, []*Connector{}, result.ops[0].connectorIn)
	assert.Equal(t,
		[][]*Connector{
			{&Connector{FromProcessor: &midiInput, FromPort: 0, ToProcessor: &osc, ToPort: 0}},
			{&Connector{FromProcessor: &midiInput, FromPort: 1, ToProcessor: &env, ToPort: 0}},
			{&Connector{FromProcessor: &midiInput, FromPort: 2, ToProcessor: &env, ToPort: 1}},
			{}, {}, {}, {},
		},
		result.ops[0].connectorOut)
	assert.Equal(t, &osc, result.ops[1].processor)
	assert.Equal(t,
		[]*Connector{
			&Connector{FromProcessor: &midiInput, FromPort: 0, ToProcessor: &osc, ToPort: 0},
		}, result.ops[1].connectorIn)
	assert.Equal(t,
		[][]*Connector{{
			&Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0},
		}}, result.ops[1].connectorOut)
	assert.Equal(t, &env, result.ops[2].processor)
	assert.Equal(t,
		[]*Connector{
			&Connector{FromProcessor: &midiInput, FromPort: 1, ToProcessor: &env, ToPort: 0},
			&Connector{FromProcessor: &midiInput, FromPort: 2, ToProcessor: &env, ToPort: 1},
		},
		result.ops[2].connectorIn)
	assert.Equal(t,
		[][]*Connector{{
			&Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1},
		}},
		result.ops[2].connectorOut)
	assert.Equal(t, &gain, result.ops[3].processor)
	assert.Equal(t,
		[]*Connector{
			&Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0},
			&Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1},
		},
		result.ops[3].connectorIn)
	assert.Equal(t,
		[][]*Connector{{
			&Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0},
		}},
		result.ops[3].connectorOut)
	assert.Equal(t, &outputTerminal, result.ops[4].processor)
	assert.Equal(t,
		[]*Connector{
			&Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0},
			&Connector{},
		},
		result.ops[4].connectorIn)
	assert.Equal(t, [][]*Connector{}, result.ops[4].connectorOut)
}
